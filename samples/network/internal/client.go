package internal

import (
	"net"
	"time"

	"github.com/yanglwd/doodle/samples/network/common"
)

type TinyClient struct {
	network    string
	localAddr  string
	remoteAddr string
	dialer     net.Dialer
	remoteConn net.Conn
	writeChan  chan []byte
	readChan   chan []byte
	closeChan  chan bool
}

func (c *TinyClient) Init(network string, localAddr string, remoteAddr string) error {
	c.writeChan = make(chan []byte, 128)
	c.readChan = make(chan []byte, 128)
	c.closeChan = make(chan bool, 8)

	localIP := net.ParseIP(localAddr)
	if localIP == nil {
		panic(common.InvalidLocalIP)
	}

	c.dialer = net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP: localIP,
		},
		Timeout: 30 * time.Second,
	}
	return nil
}

func (c *TinyClient) Serve() error {
	conn, err := c.dialer.Dial(c.network, c.remoteAddr)
	if err != nil {
		panic(err)
	}

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		panic(common.InvalidTcpConn)
	}
	tcpConn.SetKeepAlive(true)
	tcpConn.SetNoDelay(true)
	tcpConn.SetLinger(0)

	c.remoteConn = conn
mainLoop:
	for {

		select {
		case <-c.closeChan:
			break mainLoop
		default:
		}
	}

	c.remoteConn.Close()
	return nil
}

func (c *TinyClient) Stop() error {
	close(c.closeChan)
	return nil
}

func (c *TinyClient) Write(data []byte) (int, error) {
	select {
	case c.writeChan <- data:
	default:
		return -1, common.ErrSendBufferIsFull
	}
	return 0, nil
}

func (c *TinyClient) Read([]byte) error {
	return nil
}

func (c *TinyClient) writeOnce(data []byte) error {
	n, err := c.remoteConn.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return nil
	}
	return nil
}
