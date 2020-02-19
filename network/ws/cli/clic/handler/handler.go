package handler

import (
	"sync"
)

type Mode int

const (
	SYNC  Mode = 0 // 同步消息模式
	ASYNC Mode = 1 // 异步消息模式
)

//
// type Func func(data []byte)
//
// func (fn Func) Receiver(data []byte) {
// 	fn(data)
// }

type Event struct {
	clientId uint32 // 客户端Id
	receiver chan []byte
}

type Handler interface {
	EnableDebug() bool
	DebugPrintf(mid uint16, sid uint16, clientId uint32)
	Receiver(mid uint16, sid uint16, clientId uint32, data []byte)
}

type Service struct {
	sync.RWMutex
	mode   Mode
	fns    []Handler // 回调函数
	events map[uint32]*Event
}

func New(mode Mode, fn ...Handler) *Service {
	h := &Service{
		mode:   mode,
		events: make(map[uint32]*Event),
	}
	for _, f := range fn {
		h.fns = append(h.fns, f)
	}
	return h
}

func (h *Service) Fns() []Handler {
	return h.fns
}

func (h *Service) Mode() Mode {
	return h.mode
}

func (h *Service) AddSyncEvent(clientId uint32) <-chan []byte {
	h.Lock()
	defer h.Unlock()

	hevent := &Event{clientId: clientId, receiver: make(chan []byte)}
	h.events[clientId] = hevent
	return hevent.receiver
}

func (h *Service) Query(clientId uint32) *Event {
	h.Lock()
	defer h.Unlock()

	hevent, exists := h.events[clientId]
	if exists {
		return hevent
	}
	return nil
}

func (h *Service) Remove(clientId uint32) {
	h.RemoveEvent(&Event{clientId: clientId})
}

func (h *Service) RemoveEvent(removeEvent *Event) {
	h.Lock()
	defer h.Unlock()
	hevent, exists := h.events[removeEvent.ClientId()]
	if exists {
		hevent.Close()
		delete(h.events, removeEvent.ClientId())
	}
}

func (h *Service) Cancel(clientId uint32) {
	h.Remove(clientId)
}

func (h *Service) Close() {
	h.RLock()
	defer h.RUnlock()
	for _, e := range h.events {
		e.Close()
		delete(h.events, e.ClientId())
	}
}

func (h *Service) Len() int {
	h.RLock()
	defer h.RUnlock()
	return len(h.events)
}

func (e *Event) ClientId() uint32 {
	return e.clientId
}

func (e *Event) SetClientId(clientId uint32) {
	e.clientId = clientId
}

func (e *Event) Write(data []byte) {
	if e.receiver != nil {
		e.receiver <- data
	}
}

func (e *Event) Close() {
	if e.receiver != nil {
		close(e.receiver)
	}
}
