package main

import (
	. "GoStudy/gotcp"
	"GoStudy/gotcp/protocol"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type Callback struct{}

func (this *Callback) OnConnect(c *Conn) bool {
	addr := c.GetRawConn().RemoteAddr()
	c.PutExtraData(addr)
	fmt.Println("OnConnect:", addr)
	return true
}

func (this *Callback) OnMessage(c *Conn, p Packet) bool {
	lfpPacket := p.(*protocol.LfpPacket)
	fmt.Printf("OnMessage:[%v] [%v]\n", lfpPacket.GetLength(), string((lfpPacket.GetBody())))

	if bytes.Equal(lfpPacket.GetBody(), []byte("bye")) {
		fmt.Println("bye bye ", c.GetExtraData())
		return false
	}

	c.AsyncWritePacket(protocol.NewLfpPacket([]byte("echo:"+string(lfpPacket.GetBody())), false), time.Second)
	return true
}

func (this *Callback) OnClose(c *Conn) {
	fmt.Println("OnClose:", c.GetExtraData())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8989")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	config := &Config{
		AcceptTimeout:          5 * time.Second,
		ReadTimeout:            240 * time.Second,
		WriteTimeout:           240 * time.Second,
		PacketSizeLimit:        2048,
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}

	srv := NewServer(config, &Callback{}, &protocol.LfpProtocol{})

	go srv.Start(listener)
	fmt.Println("listening:", listener.Addr())

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal:", <-chSig)

	srv.Stop()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
