package tyclic

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/lw000/gocommon/network/ws/cli/clic/handler"
	"github.com/lw000/gocommon/network/ws/cli/clic/hub"
	"github.com/lw000/gocommon/network/ws/packet"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type envelope struct {
	mt  int
	msg []byte
}

var (
	TlsDialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
		ReadBufferSize:   1024 * 100,
		WriteBufferSize:  1024 * 100,
	}

	DefaultDialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		ReadBufferSize:   1024 * 100,
		WriteBufferSize:  1024 * 100,
	}
)

type WSClient struct {
	sync.RWMutex
	heartTimeSecond    int
	open               bool
	hub                *hub.Hub
	conn               *websocket.Conn
	done               chan struct{}
	input              chan *envelope
	onHeartbeatHandler func()
}

func New(heartTimeSecond int) *WSClient {
	if heartTimeSecond <= 0 {
		heartTimeSecond = 10
	}

	return &WSClient{
		hub:             hub.New(),
		done:            make(chan struct{}, 1),
		heartTimeSecond: heartTimeSecond,
		input:           make(chan *envelope, 4096),
	}
}

func (w *WSClient) Open(scheme string, host string, path string) (err error) {
	u := url.URL{Scheme: scheme, Host: host, Path: path}

	log.WithField("URL", u.String()).Info("connecting")

	var resp *http.Response
	switch scheme {
	case "ws":
		w.conn, resp, err = TlsDialer.Dial(u.String(), nil)
	case "wss":
		w.conn, resp, err = DefaultDialer.Dial(u.String(), nil)
	default:
		return errors.New("unknown ws scheme")
	}
	if err != nil {
		return err
	}

	if resp != nil {
		// log.Info(resp)
	}

	w.open = true

	go w.runControl()

	return nil
}

func (w *WSClient) RegisterAsyncMessageReceiver(mid, sid uint16, fn ...handler.Handler) {
	w.hub.RegisterAsyncMessageHandler(mid, sid, fn...)
}

func (w *WSClient) Closed() bool {
	w.RLock()
	defer w.RUnlock()
	return !w.open
}

func (w *WSClient) HandleHeartbeat(fn func()) {
	if fn == nil {
		fn = func() {}
	}
	w.onHeartbeatHandler = fn
}

func (w *WSClient) write(mt int, data []byte) error {
	w.Lock()
	defer w.Unlock()

	if err := w.conn.WriteMessage(mt, data); err != nil {
		w.open = false
		close(w.done)
		return err
	}
	return nil
}

func (w *WSClient) WriteSyncProtoMessage(mid, sid uint16, clientId uint32, pb proto.Message, waitForSeconds int) ([]byte, error) {
	if w.Closed() {
		return nil, errors.New("ws is closed")
	}

	var (
		err      error
		receiver <-chan []byte
	)

	receiver = w.hub.RegisterSyncMessageHandler(mid, sid, clientId)

	var buf []byte
	if pb != nil {
		buf, err = proto.Marshal(pb)
		if err != nil {
			return nil, err
		}
	}

	packet := typacket.NewPacket(mid, sid, clientId)
	if err = packet.Encode(buf); err != nil {
		return nil, err
	}

	select {
	case w.input <- &envelope{mt: websocket.BinaryMessage, msg: packet.Data()}:
	default:
		return nil, errors.New("message buffer is full")
	}

	// TODO: 等待超时取消网络请求任务
	var data []byte
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(waitForSeconds))
	select {
	case data = <-receiver:
	case <-ctx.Done():
		w.hub.CancelSyncMessageHandler(mid, sid, clientId)
		return nil, errors.New(fmt.Sprintf("wait message timeout.[%ds]", waitForSeconds))
	}
	return data, err
}

func (w *WSClient) WriteAsyncProtoMessage(mid, sid uint16, clientId uint32, pb proto.Message) error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	var (
		err error
		buf []byte
	)

	if pb != nil {
		buf, err = proto.Marshal(pb)
		if err != nil {
			return err
		}
	}

	packet := typacket.NewPacket(mid, sid, clientId)
	if err = packet.Encode(buf); err != nil {
		return err
	}

	select {
	case w.input <- &envelope{mt: websocket.BinaryMessage, msg: packet.Data()}:
	default:
		return errors.New("message buffer is full")
	}

	return nil
}

func (w *WSClient) runControl() {
	defer func() {
		if x := recover(); x != nil {

		}
		log.Error("ws control exit")
	}()

	t := time.NewTicker(time.Second * time.Duration(w.heartTimeSecond))
	defer t.Stop()

loop:
	for {
		select {
		case <-t.C: // TODO:心跳处理
			go w.onHeartbeatHandler()
		case msg := <-w.input: // TODO: 发送网络发送数据
			if err := w.conn.WriteMessage(msg.mt, msg.msg); err != nil {
				w.open = false
				w.done <- struct{}{}
				log.Info(err)
			}
		case <-w.done:
			w.Lock()
			if err := w.conn.Close(); err != nil {
				log.Error(err)
			}
			w.open = false
			w.hub.Close()
			w.Unlock()
			break loop
		}
	}
}

func (w *WSClient) Run() error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	go w.readMessage()

	return nil
}

func (w *WSClient) readMessage() {
	defer func() {
		log.Error("ws read exit")
	}()

	for {
		mt, message, err := w.conn.ReadMessage()
		if err != nil {
			w.open = false
			w.done <- struct{}{}
			return
		}

		if mt != websocket.BinaryMessage {
			log.Error("protocol error")
			return
		}

		if err = w.hub.DispatchMessage(message); err != nil {
			log.Error(err)
		}

		ln := w.hub.Len()
		if ln > 0 {
			log.WithFields(log.Fields{"len": ln}).Warning("ws消息分发")
		}
	}
}

func (w *WSClient) Stop() {
	if w == nil {
		return
	}
	w.done <- struct{}{}
}
