package testhelper

import (
	"GoStudy/loadgen/lib"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

const (
	DELIM = '\n'
)

type TcpComm struct {
	addr string
}

func NewTcpComm(addr string) *TcpComm {
	return &TcpComm{addr: addr}
}

func (comm *TcpComm) BuildReq() lib.RawReq {
	id := time.Now().UnixNano()
	sreq := ServerReq{
		Id:       id,
		Operands: []int{int(rand.Int31()), int(rand.Int31())},
		Operator: func() string {
			op := []string{"+", "-", "*", "/"}
			return op[rand.Int31n(100)%4]
		}(),
	}
	bytes, err := json.Marshal(sreq)
	if err != nil {
		panic(err)
	}
	rawReq := lib.RawReq{Id: id, Req: bytes}
	return rawReq
}

func (comm *TcpComm) Call(req []byte, timeoutNs time.Duration) ([]byte, error) {
	conn, err := net.DialTimeout("tcp", comm.addr, timeoutNs)
	if err != nil {
		return nil, err
	}
	_, err = write(conn, req, DELIM)
	if err != nil {
		return nil, err
	}
	return read(conn, DELIM)
}

func (comm *TcpComm) CheckResp(req lib.RawReq, resp lib.RawResp) *lib.CallResult {
	var callResult lib.CallResult
	callResult.Id = resp.Id
	callResult.Req = req
	callResult.Resp = resp
	var sreq ServerReq
	err := json.Unmarshal(req.Req, &sreq)
	if err != nil {
		callResult.Code = lib.RESULT_CODE_FATAL_CALL
		callResult.Msg = fmt.Sprintf("Incorrectly formatted Req: %s!\n", string(req.Req))
		return &callResult
	}

	var sresp ServerResp
	err = json.Unmarshal(resp.Resp, &sresp)
	if err != nil {
		callResult.Code = lib.RESULT_CODE_ERROR_RESPONSE
		callResult.Msg = fmt.Sprintf("Incorrectly formatted Resp: %s!\n", string(resp.Resp))
		return &callResult
	}
	if sresp.Id != sreq.Id {
		callResult.Code = lib.RESULT_CODE_ERROR_RESPONSE
		callResult.Msg = fmt.Sprintf("Inconsistent raw id! (%d != %d)\n", req.Id, resp.Id)
		return &callResult
	}
	if sresp.Err != nil {
		callResult.Code = lib.RESULT_CODE_ERROR_CALEE
		callResult.Msg =
			fmt.Sprintf("Abnormal server: %s!\n", sresp.Err)
		return &callResult
	}
	if sresp.Result != op(sreq.Operands, sreq.Operator) {
		callResult.Code = lib.RESULT_CODE_ERROR_RESPONSE
		callResult.Msg =
			fmt.Sprintf("Incorrect result: %s!\n", genFormula(sreq.Operands, sreq.Operator, sresp.Result, false))
		return &callResult
	}
	callResult.Code = lib.RESULT_CODE_SUCCESS
	callResult.Msg = fmt.Sprintf("Success. (%s)", sresp.Formula)
	return &callResult
}

func read(conn net.Conn, delim byte) ([]byte, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return nil, err
		}
		readByte := readBytes[0]
		if readByte == delim {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.Bytes(), nil
}

func write(conn net.Conn, data []byte, delim byte) (int, error) {
	writer := bufio.NewWriter(conn)
	n, err := writer.Write(data)
	if err == nil {
		writer.WriteByte(delim)
	}
	if err == nil {
		writer.Flush()
	}
	return n, err
}
