package tyhub

import (
	"errors"
	"github.com/gorilla/websocket"
	"gocommon/network/ws/packet"
	"sync"
)

var ErrNotFound = errors.New("hub: not found handler")
var ErrProto = errors.New("hub: proto error")

type hubKey struct {
	mid uint16
	sid uint16
}

type Hub struct {
	handlers sync.Map
}

func New() *Hub {
	return &Hub{}
}

func (h *Hub) RegisterHandler(mid, sid uint16, handler Handler) {
	k := hubKey{mid: mid, sid: sid}
	_, ok := h.handlers.Load(k)
	if !ok {
		h.handlers.Store(k, handler)
	}
}

func (h *Hub) UnregisterHandler(mid, sid uint16) {
	key := hubKey{mid: mid, sid: sid}
	_, ok := h.handlers.Load(key)
	if ok {
		h.handlers.Delete(key)
	}
}

func (h *Hub) Query(mid, sid uint16) Handler {
	key := hubKey{mid: mid, sid: sid}
	if v, ok := h.handlers.Load(key); ok {
		return v.(Handler)
	}
	return nil
}

func (h *Hub) Close() {

}

func (h *Hub) DispatchMessage(conn *websocket.Conn, msg []byte) error {
	pk, err := typacket.NewPacketWithData(msg)
	if err != nil {
		return ErrProto
	}

	handler := h.Query(pk.Mid(), pk.Sid())
	if handler == nil {
		return ErrNotFound
	}

	handler.Receiver(conn, pk)

	return nil
}
