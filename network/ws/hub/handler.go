package tyhub

import (
	"github.com/gorilla/websocket"
	"gocommon/network/ws/packet"
)

type receiverFunc func(conn *websocket.Conn, pk *typacket.Packet)

func (fn receiverFunc) Receiver(conn *websocket.Conn, pk *typacket.Packet) {
	fn(conn, pk)
}

type Handler interface {
	Receiver(conn *websocket.Conn, pk *typacket.Packet)
}
