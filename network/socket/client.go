package tysockt

import (
	"github.com/labstack/gommon/log"
	"net"
)

type Client struct {
	conn      net.Conn
	connected bool
	quit      chan struct{}
	onMessage func(data []byte) error
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) OnMessage(onMessage func(data []byte) error) {
	c.onMessage = onMessage
}

func (c *Client) Connected() bool {
	return c.connected
}

func (c *Client) Open(host string, port string) error {
	var err error
	c.conn, err = net.Dial("tcp", net.JoinHostPort(host, port))

	checkError(err)

	c.connected = true

	go c.run()

	return nil
}

func (c *Client) Send(data []byte) error {
	n, err := c.conn.Write(data)
	if err != nil {
		log.Error("connected closed")
	}

	if n > 0 {

	}

	return nil
}

func (c *Client) run() {
	var (
		n   int
		err error
	)
	buf := make([]byte, 1024)
	for {
		n, err = c.conn.Read(buf)
		if err != nil {
			break
		}

		if n > 0 {

		}

		if c.onMessage != nil {
			err = c.onMessage(buf[0:n])
		}
	}
}

func (c *Client) Close() error {
	err := c.conn.Close()
	if err != nil {

	}
	return nil
}
