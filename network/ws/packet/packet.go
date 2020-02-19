package typacket

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Packet struct {
	ver       uint8
	mid       uint16
	sid       uint16
	checkCode uint32
	clientId  uint32
	data      []byte
}

// var (
// 	pool *sync.Pool
// )
//
// func init() {
// 	pool = &sync.Pool{New: func() interface{} {
// 		return &Packet{}
// 	}}
// }

func NewPacket(mid, sid uint16, clientId uint32) *Packet {
	p := &Packet{
		mid:      mid,
		sid:      sid,
		clientId: clientId,
	}
	return p
}

func NewPacketWithData(data []byte) (*Packet, error) {
	if len(data) == 0 {
		return nil, errors.New("data item is zero")
	}
	p := &Packet{}
	buf := bytes.NewBuffer(data)
	if err := p.readHead(buf); err != nil {
		return nil, err
	}
	p.data = buf.Bytes()
	return p, nil
}

func (p *Packet) writeHead(buf *bytes.Buffer) (err error) {
	if err = binary.Write(buf, binary.LittleEndian, p.ver); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.checkCode); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.mid); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.sid); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.clientId); err != nil {
		return err
	}
	return err
}

func (p *Packet) readHead(buf *bytes.Buffer) (err error) {
	if err = binary.Read(buf, binary.LittleEndian, &p.ver); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.checkCode); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.mid); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.sid); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.clientId); err != nil {
		return err
	}
	return err
}

// Encode 编码数据包
func (p *Packet) Encode(data []byte) error {
	buf := &bytes.Buffer{}
	err := p.writeHead(buf)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		var n int
		n, err = buf.Write(data)
		if err != nil {
			return err
		}
		if n < 0 {

		}
	}

	p.data = buf.Bytes()

	return nil
}

// EncodeProto 编码数据包
// func (p *Packet) EncodeProto(pb proto.Message) error {
// 	if pb == nil {
// 		err := p.Encode([]byte{})
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		data, err := proto.Marshal(pb)
// 		if err != nil {
// 			return err
// 		}
// 		err = p.Encode(data)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (p Packet) Ver() uint8 {
	return p.ver
}

func (p *Packet) CheckCode() uint32 {
	return p.checkCode
}

func (p Packet) Mid() uint16 {
	return p.mid
}

func (p Packet) Sid() uint16 {
	return p.sid
}

func (p *Packet) ClientId() uint32 {
	return p.clientId
}

func (p Packet) Data() []byte {
	return p.data
}

func (p *Packet) SetClientId(clientId uint32) {
	p.clientId = clientId
}

func (p *Packet) SetCheckCode(checkCode uint32) {
	p.checkCode = checkCode
}

func (p Packet) String() string {
	return fmt.Sprintf("{ver:%d mid:%d sid:%d checkCode:%d clientId:%d datalen:%d}", p.ver, p.mid, p.sid, p.checkCode, p.clientId, len(p.data))
}
