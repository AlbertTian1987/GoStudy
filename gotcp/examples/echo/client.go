package main

import (
	"GoStudy/gotcp/protocol"
	"fmt"
	"log"
	"net"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8989")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	LfpProtocol := &protocol.LfpProtocol{}

	data := make([]byte, 1024)
	for {
		fmt.Scan(&data)
		if string(data) == "exit" {
			break
		}
		conn.Write(protocol.NewLfpPacket(data, false).Serialize())
		p, err := LfpProtocol.ReadPacket(conn, 1024)
		if err == nil {
			lfpPacket := p.(*protocol.LfpPacket)
			fmt.Printf("Server reply:[%v] [%v]\n", lfpPacket.GetLength(), string(lfpPacket.GetBody()))
		}
	}

	conn.Write(protocol.NewLfpPacket([]byte("bye"), false).Serialize())
	conn.Close()

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
