package gotcp

import (
	"net"
	"sync"
	"time"
)

type Server struct {
	basic *basicSrv
}

type Config struct {
	AcceptTimeout          time.Duration
	ReadTimeout            time.Duration
	WriteTimeout           time.Duration
	PacketSizeLimit        uint32
	PacketSendChanLimit    uint32
	PacketReceiveChanLimit uint32
}

type basicSrv struct {
	config    *Config
	callback  ConnCallback
	protocol  Protocol
	exitChan  chan struct{}
	waitGroup *sync.WaitGroup
}

func newBasicSrv(config *Config, callback ConnCallback, protocol Protocol) *basicSrv {
	return &basicSrv{
		config:    config,
		callback:  callback,
		protocol:  protocol,
		exitChan:  make(chan struct{}),
		waitGroup: &sync.WaitGroup{},
	}
}

func NewServer(config *Config, callback ConnCallback, protocol Protocol) *Server {
	basicSrv := newBasicSrv(config, callback, protocol)
	return &Server{basicSrv}
}

func (s *Server) Start(listener *net.TCPListener) {
	s.basic.waitGroup.Add(1)
	defer func() {
		listener.Close()
		s.basic.waitGroup.Done()
	}()

	for {
		select {
		case <-s.basic.exitChan:
			return
		default:
		}

		listener.SetDeadline(time.Now().Add(s.basic.config.AcceptTimeout))

		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		go newConn(conn, s.basic).Do()
	}
}

func (s *Server) Stop() {
	close(s.basic.exitChan)
	s.basic.waitGroup.Wait()
}

func (s *Server) Dial(network, address string, config *Config, callback ConnCallback, protocol Protocol) (*Conn, error) {
	tcpAddr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	basic := newBasicSrv(config, callback, protocol)
	return newConn(conn, basic), nil
}
