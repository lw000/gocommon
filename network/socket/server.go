package tysockt

import (
	"net"
	"time"

	"log"
)

type Server struct {
}

func handleClient(conn net.Conn) {
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

func (s *Server) Run(addr string) error {
	listen, err := net.Listen("tcp", addr)
	checkError(err)

	defer func() {
		err = listen.Close()
		if err != nil {

		}
	}()

	log.Printf("server start... port:[%s]", addr)

	for {
		var conn net.Conn
		conn, err = listen.Accept()
		if err != nil {
			continue
		}

		log.Printf("[%s]", conn.RemoteAddr().String())

		go handleClient(conn)
	}
}
