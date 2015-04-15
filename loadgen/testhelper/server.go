package testhelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"
)

type ServerReq struct {
	Id       int64
	Operands []int
	Operator string
}

type ServerResp struct {
	Id      int64
	Formula string
	Result  int
	Err     error
}

func op(operands []int, operator string) int {
	var result int
	switch operator {
	case "+":
		for _, v := range operands {
			if result == 0 {
				result = v
			} else {
				result += v
			}
		}
	case "-":
		for _, v := range operands {
			if result == 0 {
				result = v
			} else {
				result -= v
			}
		}
	case "*":
		for _, v := range operands {
			if result == 0 {
				result = v
			} else {
				result *= v
			}
		}
	case "/":
		for _, v := range operands {
			if result == 0 {
				result = v
			} else {
				result /= v
			}
		}
	}
	return result

}

func genFormula(operands []int, operator string, result int, equal bool) string {
	var buff bytes.Buffer
	n := len(operands)
	for i := 0; i < n; i++ {
		if i > 0 {
			buff.WriteString(" ")
			buff.WriteString(operator)
			buff.WriteString(" ")
		}

		buff.WriteString(strconv.Itoa(operands[i]))
	}
	if equal {
		buff.WriteString(" = ")
	} else {
		buff.WriteString(" != ")
	}
	buff.WriteString(strconv.Itoa(result))
	return buff.String()
}

func reqHandler(conn net.Conn) {
	var errMsg string
	var sresp ServerResp
	req, err := read(conn, DELIM)
	if err != nil {
		errMsg = fmt.Sprintf("Server: Req Read Error : %s\n", err)
	} else {
		var sreq ServerReq
		err := json.Unmarshal(req, &sreq)
		if err != nil {
			errMsg = fmt.Sprintf("Server: Req Unmarshal Error: %s", err)
		} else {
			sresp.Id = sreq.Id
			sresp.Result = op(sreq.Operands, sreq.Operator)
			sresp.Formula = genFormula(sreq.Operands, sreq.Operator, sresp.Result, true)
		}
	}
	if errMsg != "" {
		sresp.Err = errors.New(errMsg)
	}
	bytes, err := json.Marshal(sresp)
	if err != nil {
		fmt.Errorf("Server: Resp Marshal Error: %s", err)
	}
	_, err = write(conn, bytes, DELIM)
	if err != nil {
		fmt.Errorf("Server: Resp Write error: %s", err)
	}
}

type TcpServer struct {
	listener net.Listener
	active   bool
	lock     *sync.Mutex
}

func (this *TcpServer) init(addr string) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.active {
		return nil
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	this.listener = ln
	this.active = true
	return nil
}

func (this *TcpServer) Listen(addr string) error {
	err := this.init(addr)
	if err != nil {
		return err
	}
	go func(active *bool) {
		for {
			conn, err := this.listener.Accept()
			if err != nil {
				fmt.Printf("Server: Request Acception Error: %s\n", err)
				continue
			}
			go reqHandler(conn)
			runtime.Gosched()
		}
	}(&this.active)
	return nil
}

func (this *TcpServer) Close() bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.active {
		this.listener.Close()
		this.active = false
		return true
	} else {
		return false
	}
}

func NewTcpServer() *TcpServer {
	return &TcpServer{lock: new(sync.Mutex)}
}
