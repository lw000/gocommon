package tyclis

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/lw000/gocommon/network/ws/hub"
	"github.com/lw000/gocommon/network/ws/packet"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var (
	TlsDialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
)

type WSSClient struct {
	sync.RWMutex
	conn            *websocket.Conn
	hub             *tyhub.Hub
	done            chan struct{}
	heartTimeSecond int
	open            bool
	heartBeatFn     func() error
}

func New(heartTimeSecond int) *WSSClient {
	if heartTimeSecond < 0 {
		heartTimeSecond = 0
	}

	return &WSSClient{
		hub:             tyhub.New(),
		done:            make(chan struct{}, 1),
		heartTimeSecond: heartTimeSecond,
	}
}

func (w *WSSClient) Closed() bool {
	w.Lock()
	defer w.Unlock()
	return !w.open
}

func (w *WSSClient) Open(scheme string, host string, path string) (err error) {
	u := url.URL{Scheme: scheme, Host: host, Path: path}

	log.Println("connecting to ", u.String())

	if scheme == "wss" {
		w.conn, _, err = TlsDialer.Dial(u.String(), nil)
	} else if scheme == "ws" {
		w.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	} else {
		return errors.New(fmt.Sprintf("unknow scheme [%s]", scheme))
	}

	if err != nil {
		return err
	}

	w.open = true

	go w.runControl()

	return nil
}

func (w *WSSClient) HanldeHeartbeat(fn func() error) {
	if fn == nil {
		fn = func() error { return nil }
	}
	w.heartBeatFn = fn
}

func (w *WSSClient) write(mt int, buf []byte) error {
	w.Lock()
	defer w.Unlock()
	if err := w.conn.WriteMessage(mt, buf); err != nil {
		w.open = false
		w.done <- struct{}{}
		return errors.New("error")
	}

	return nil
}

func (w *WSSClient) AddHandler(mid, sid uint16, handler tyhub.Handler) {
	w.hub.RegisterHandler(mid, sid, handler)
}

func (w *WSSClient) RemoveHandler(mid, sid uint16) {
	w.hub.UnregisterHandler(mid, sid)
}

func (w *WSSClient) WriteBinaryMessage(mid, sid uint16, clientId uint32, data []byte) error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	pk := typacket.NewPacket(mid, sid, clientId)
	var err error
	if err = pk.Encode(data); err != nil {
		return err
	}

	if err = w.write(websocket.BinaryMessage, pk.Data()); err != nil {
		return err
	}

	return nil
}

func (w *WSSClient) WriteProtoMessage(mid, sid uint16, clientId uint32, pb proto.Message) error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	pk := typacket.NewPacket(mid, sid, clientId)
	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	if err = pk.Encode(data); err != nil {
		return err
	}

	if err = w.write(websocket.BinaryMessage, pk.Data()); err != nil {
		return err
	}

	return nil
}

func (w *WSSClient) Run() error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	go w.readMessage()

	return nil
}

func (w *WSSClient) Stop() {
	w.done <- struct{}{}
}

func (w *WSSClient) Close() {
	w.hub.Close()
}

func (w *WSSClient) runControl() {
	defer func() {
		log.Error("ws control exit")
	}()

	heartTicker := time.NewTicker(time.Second * time.Duration(w.heartTimeSecond))
	defer heartTicker.Stop()

loop:
	for {
		select {
		case <-heartTicker.C: // TODO:心跳处理
			err := w.heartBeatFn()
			// er := w.send(websocket.PingMessage, []byte{})
			if err != nil {
			}
			break
		case <-w.done:
			w.Lock()
			err := w.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
			}

			if err = w.conn.Close(); err != nil {
			}

			w.open = false
			w.Unlock()
			break loop
		}
	}
}

func (w *WSSClient) readMessage() {
	defer func() {
		log.Error("ws readMessage exit")
	}()

	for {
		mt, message, err := w.conn.ReadMessage()
		if err != nil {
			log.Error(err)
			w.done <- struct{}{}
			return
		}

		if mt != websocket.BinaryMessage {
			return
		}

		if err = w.hub.DispatchMessage(w.conn, message); err != nil {
			log.Error(err)
		}
	}
}
