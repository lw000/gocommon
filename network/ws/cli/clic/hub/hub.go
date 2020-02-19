package hub

import (
	"errors"
	"fmt"
	"github.com/lw000/gocommon/network/ws/cli/clic/handler"
	"github.com/lw000/gocommon/network/ws/packet"
	"sync"
)

type Key struct {
	mid uint16
	sid uint16
}

type Hub struct {
	handlers sync.Map
}

func New() *Hub {
	return &Hub{}
}

// RegisterSyncHandler 注册同步消息handler
func (h *Hub) RegisterSyncMessageHandler(mid, sid uint16, clientId uint32) <-chan []byte {
	var hmgr *handler.Service
	key := Key{mid: mid, sid: sid}
	value, exists := h.handlers.Load(key)
	if exists {
		hmgr = value.(*handler.Service)
	} else {
		hmgr = handler.New(handler.SYNC)
		h.handlers.Store(key, hmgr)
	}
	return hmgr.AddSyncEvent(clientId)
}

// RegisterSyncHandler 注册异步消息handler
func (h *Hub) RegisterAsyncMessageHandler(mid, sid uint16, fn ...handler.Handler) {
	var hmgr *handler.Service
	key := Key{mid: mid, sid: sid}
	value, exists := h.handlers.Load(key)
	if exists {
		hmgr = value.(*handler.Service)
	} else {
		hmgr = handler.New(handler.ASYNC, fn...)
		h.handlers.Store(key, hmgr)
	}
}

// CancelSyncHandler  取消同步消息handler
func (h *Hub) CancelSyncMessageHandler(mid, sid uint16, clientId uint32) {
	key := Key{mid: mid, sid: sid}
	if value, exists := h.handlers.Load(key); exists {
		value.(*handler.Service).Cancel(clientId)
	}
}

func (h *Hub) Query(mid, sid uint16) *handler.Service {
	key := Key{mid: mid, sid: sid}
	if value, exists := h.handlers.Load(key); exists {
		return value.(*handler.Service)
	}
	return nil
}

func (h *Hub) Len() int {
	var ln int
	h.handlers.Range(func(key interface{}, value interface{}) bool {
		ln += value.(*handler.Service).Len()
		return false
	})
	return ln
}

func (h *Hub) DispatchMessage(msg []byte) error {
	if len(msg) == 0 {
		return errors.New("read ws msg is empty")
	}

	packet, err := typacket.NewPacketWithData(msg)
	if err != nil {
		return err
	}

	service := h.Query(packet.Mid(), packet.Sid())
	if service == nil {
		var ErrNotFound = errors.New(fmt.Sprintf("hub: not found handler [mid=%x, sid=%x]", packet.Mid(), packet.Sid()))
		return ErrNotFound
	}

	switch service.Mode() {
	case handler.SYNC:
		event := service.Query(packet.ClientId())
		if event == nil {
			return fmt.Errorf("ws request was cancelled due to timeout. detail: %+v", packet)
		}
		event.Write(packet.Data())
		service.RemoveEvent(event)
	case handler.ASYNC:
		for _, fn := range service.Fns() {
			fnCopy := fn
			go func() {
				defer func() {
					if x := recover(); x != nil {
					}
				}()
				if fnCopy.EnableDebug() {
					fnCopy.DebugPrintf(packet.Mid(), packet.Sid(), packet.ClientId())
				}
				fnCopy.Receiver(packet.Mid(), packet.Sid(), packet.ClientId(), packet.Data())
			}()
		}
	default:
	}

	return nil
}

func (h *Hub) Close() {
	h.handlers.Range(func(key interface{}, value interface{}) bool {
		value.(*handler.Service).Close()
		h.handlers.Delete(key)
		return true
	})
}
