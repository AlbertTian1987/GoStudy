package gotcp

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrConnClosing   = errors.New("use closed network connection")
	ErrWriteBlocking = errors.New("write packet was blocking")
	ErrReadBlocking  = errors.New("read packet was blocking")
)

type Conn struct {
	basic             *basicSrv
	conn              *net.TCPConn
	extraData         interface{}
	closeOnce         sync.Once
	closeFlag         int32
	closeChan         chan struct{}
	packetSendChan    chan Packet
	packetReceiveChan chan Packet
}

type ConnCallback interface {
	// OnConnect is called when the connection was accepted,
	// If the return value of false is closed
	OnConnect(*Conn) bool
	// OnMessage is called when the connection receives a packet,
	// If the return value of false is closed
	OnMessage(*Conn, Packet) bool
	// OnClose is called when the connection closed
	OnClose(*Conn)
}

func newConn(conn *net.TCPConn, basic *basicSrv) *Conn {
	return &Conn{
		basic:             basic,
		conn:              conn,
		closeChan:         make(chan struct{}),
		packetSendChan:    make(chan Packet, basic.config.PacketSendChanLimit),
		packetReceiveChan: make(chan Packet, basic.config.PacketReceiveChanLimit),
	}
}

func (c *Conn) PutExtraData(data interface{}) {
	c.extraData = data
}

func (c *Conn) GetExtraData() interface{} {
	return c.extraData
}

func (c *Conn) GetRawConn() *net.TCPConn {
	return c.conn
}

func (c *Conn) IsClosed() bool {
	return atomic.LoadInt32(&c.closeFlag) == 1
}

func (c *Conn) AsyncReadPacket(timeout time.Duration) (Packet, error) {
	if c.IsClosed() {
		return nil, ErrConnClosing
	}
	if timeout == 0 {
		select {
		case p := <-c.packetReceiveChan:
			return p, nil
		default:
			return nil, ErrReadBlocking
		}
	} else {
		select {
		case p := <-c.packetReceiveChan:
			return p, nil
		case <-c.closeChan:
			return nil, ErrConnClosing
		case <-time.After(timeout):
			return nil, ErrReadBlocking
		}
	}
}

func (c *Conn) AsyncWritePacket(p Packet, timeout time.Duration) error {
	if c.IsClosed() {
		return ErrConnClosing
	}
	if timeout == 0 {
		select {
		case c.packetSendChan <- p:
			return nil
		default:
			return ErrWriteBlocking
		}
	} else {
		select {
		case c.packetSendChan <- p:
			return nil
		case <-c.closeChan:
			return ErrConnClosing
		case <-time.After(timeout):
			return ErrWriteBlocking
		}
	}

}

func (c *Conn) Do() {
	if !c.basic.callback.OnConnect(c) {
		return
	}

	c.basic.waitGroup.Add(3)
	go c.handleLoop()
	go c.readLoop()
	go c.writeLoop()
}

func (c *Conn) Close() {
	c.closeOnce.Do(func() {
		atomic.StoreInt32(&c.closeFlag, 1)
		close(c.closeChan)
		c.conn.Close()
		c.basic.callback.OnClose(c)
	})
}

func (c *Conn) readLoop() {
	defer func() {
		recover()
		c.Close()
		c.basic.waitGroup.Done()
	}()

	for {
		select {
		case <-c.basic.exitChan:
			return
		case <-c.closeChan:
			return
		default:
		}

		c.conn.SetReadDeadline(time.Now().Add(c.basic.config.ReadTimeout))
		recPacket, err := c.basic.protocol.ReadPacket(c.conn, c.basic.config.PacketSizeLimit)
		if err != nil {
			return
		}
		c.packetReceiveChan <- recPacket

	}
}

func (c *Conn) writeLoop() {
	defer func() {
		recover()
		c.Close()
		c.basic.waitGroup.Done()
	}()

	for {
		select {
		case <-c.basic.exitChan:
			return
		case <-c.closeChan:
			return
		case p := <-c.packetSendChan:
			c.conn.SetWriteDeadline(time.Now().Add(c.basic.config.WriteTimeout))
			if _, err := c.conn.Write(p.Serialize()); err != nil {
				return
			}
		}
	}
}

func (c *Conn) handleLoop() {
	defer func() {
		recover()
		c.Close()
		c.basic.waitGroup.Done()
	}()

	for {
		select {
		case <-c.basic.exitChan:
			return
		case <-c.closeChan:
			return
		case p := <-c.packetReceiveChan:
			if !c.basic.callback.OnMessage(c, p) {
				return
			}
		}
	}
}
