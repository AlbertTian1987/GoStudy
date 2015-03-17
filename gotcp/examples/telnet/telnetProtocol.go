package telnet

import (
	. "GoStudy/gotcp"
	"bytes"
	"fmt"
	"io"
	"strings"
)

var (
	endTag = []byte("\r\n")
)

type TelnetPacket struct {
	pLen  uint32
	pType string
	pData []byte
}

func (p *TelnetPacket) Serialize() []byte {
	buf := p.pData
	buf = append(buf, endTag...)
	return buf
}

func (p *TelnetPacket) GetLen() uint32 {
	return p.pLen
}

func (p *TelnetPacket) GetType() string {
	return p.pType
}

func (p *TelnetPacket) GetData() []byte {
	return p.pData
}

func NewTelnetPacket(pType string, pData []byte) *TelnetPacket {
	return &TelnetPacket{
		pLen:  uint32(len(pData)),
		pType: pType,
		pData: pData,
	}
}

type TelnetProtocol struct {
}

func (this *TelnetProtocol) ReadPocket(r io.Reader, packetSizeLimit uint32) (Packet, error) {
	fullBuf := bytes.NewBuffer([]byte{})
	for {
		data := make([]byte, packetSizeLimit)
		readLength, err := r.Read(data)
		if err != nil {
			return nil, err
		}
		if readLength == 0 {
			return nil, ErrConnClosing
		}

		fullBuf.Write(data[:readLength])
		index := bytes.Index(fullBuf.Bytes(), endTag)
		if index > -1 {
			command := fullBuf.Next(index)
			fullBuf.Next(2)

			commandList := strings.Split(string(command), " ")
			if len(commandList) > 1 {
				return NewTelnetPacket(commandList[0], []byte(commandList[1])), nil
			} else {
				if commandList[0] == "quit" {
					return NewTelnetPacket("quit", command), nil
				} else {
					return NewTelnetPacket("unknow", command), nil
				}
			}
		}

	}
}

type TelnetCallback struct {
	connectCount int
	closeCount   int
	messageCount int
}

func (this *TelnetCallback) OnConnect(c *Conn) bool {
	this.connectCount++
	c.PutExtraData(this.connectCount)
	fmt.Printf("OnConnect[%s][***%v***]\n", c.GetRawConn().RemoteAddr(), c.GetExtraData().(int))

	c.AsyncWritePacket(NewTelnetPacket("unknow", []byte("Welcome to this Telnet server")), 0)
	return true
}

func (this *TelnetCallback) OnMessage(c Conn, p Packet) bool {
	packet := p.(*TelnetPacket)

	fmt.Printf("OnMessage[%s][***%v***]:[%v]\n", c.GetRawConn().RemoteAddr(),
		c.GetExtraData().(int), string(packet.GetData()))
	this.messageCount++
	command := packet.GetData()
	commandType := packet.GetType()
	switch commandType {
	case "echo":
		c.AsyncWritePacket(NewTelnetPacket("echo", command), 0)
	case "login":
		c.AsyncWritePacket(NewTelnetPacket("login", []byte(string(command)+" has login")), 0)
	case "quit":
		return false
	default:
		c.AsyncWritePacket(NewTelnetPacket("unknow", []byte("unknow command")), 0)
	}
	return true
}

func (this *TelnetCallback) OnClose(c Conn) {
	this.closeCount++
	fmt.Printf("OnClose[%s][***%v***]\n", c.GetRawConn().RemoteAddr(), c.GetExtraData().(int))
}
