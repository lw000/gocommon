package tysockt

import (
	"context"
	"net"
	"time"

	"log"
)

type Server struct {
	listen     net.Listener
	cancelFunc context.CancelFunc
}

func NewServer() *Server {
	return &Server{}
}

func handleClient(ctx context.Context, conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {

		}
	}()

	var (
		n   int
		err error
		buf []byte
	)

	buf = make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err = conn.Read(buf)
			if err != nil {
				log.Println("connected closed")
				break
			}
			if n > 0 {

			}

			log.Printf("read:%s", string(buf[0:n]))

			n, err = conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
			if err != nil {
				log.Println("connected closed")
				break
			}

			if n > 0 {

			}
		}
	}
}

func (s *Server) Run(addr string) error {
	var err error
	s.listen, err = net.Listen("tcp", addr)
	if err != nil {
		log.Printf("listen error:%s", err.Error())
		return err
	}
	log.Printf("server start... port:[%s]", addr)

	var ctx context.Context
	ctx, s.cancelFunc = context.WithCancel(context.Background())
	for {
		var conn net.Conn
		conn, err = s.listen.Accept()
		if err != nil {
			log.Printf("accept error: %s", err.Error())
			continue
		}
		log.Printf("[%s]", conn.RemoteAddr().String())

		go handleClient(ctx, conn)
	}
}

func (s *Server) Stop() {
	var err error
	err = s.listen.Close()
	if err != nil {

	}
	s.cancelFunc()
}
