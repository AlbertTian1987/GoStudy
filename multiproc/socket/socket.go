package main
import (
    "net"
    "fmt"
    "time"
    "bufio"
    "io"
    "strconv"
    "bytes"
    "math"
    "math/rand"
    "sync"
    "errors"
)




const(
    SERVER_NETWORK = "tcp"
    SERVER_ADDRESS = "127.0.0.1:8085"
    DELIMITER = '\t'
)
var wg sync.WaitGroup


func main() {
    wg.Add(2)
    go serverGo()
    time.Sleep(500 * time.Millisecond)
    go clientGo(123)
    wg.Wait()
}

func serverGo() {
    var listener net.Listener
    listener,err:=net.Listen(SERVER_NETWORK,SERVER_ADDRESS)
    if err!=nil {
        printLog("Listener error :",err)
        return
    }

    defer listener.Close()

    for {
        conn,err:=listener.Accept()
        if err!=nil {
            printLog("Connect Error:",err)
            continue
        }
        printLog("Connect from ",conn.RemoteAddr().String())
        go handelConn(conn)
    }

}

func handelConn(conn net.Conn){
    defer func(){
        conn.Close()
        wg.Done()
    }()
    for {
        conn.SetReadDeadline(time.Now().Add(10*time.Second))

        reader:=bufio.NewReader(conn)
        data,err:=reader.ReadBytes(DELIMITER)

        if err!=nil{
            if err==io.EOF {
                printLog("the connect is closed")
            }else {
                printLog("read error :",err)
            }
            break
        }

        req32,err:=convertToInt32(string(data))

        if err!=nil {
            _,err=writeTo(conn,err.Error())
            if err!=nil {
                printLog("Server Write error:",err)
            }
            continue
        }

        resp64:=math.Sqrt(float64(req32))
        respMsg := fmt.Sprintf("the %d sqrt is %f",req32,resp64)

        _,err=writeTo(conn,respMsg)
        if err!=nil {
            printLog("Server Write error:",err.Error())
        }

        printLog("Sent response ",respMsg)
    }
}

func convertToInt32(str string) (int32, error) {
    num, err := strconv.Atoi(str)
    if err != nil {
        printLog(fmt.Sprintf("Parse Error: %s\n", err))
        return 0, errors.New(fmt.Sprintf("'%s' is not integer!", str))
    }
    if num > math.MaxInt32 || num < math.MinInt32 {
        printLog(fmt.Sprintf("Convert Error: The integer %s is too large/small.\n", num))
        return 0, errors.New(fmt.Sprintf("'%s' is not 32-bit integer!", num))
    }
    return int32(num), nil
}

func writeTo(conn net.Conn,str string)(int,error){
    var buff bytes.Buffer
    buff.WriteString(str)
    buff.WriteByte(DELIMITER)
    return conn.Write(buff.Bytes())
}


func printLog(logs ...interface{}){
    fmt.Println(logs...)
}


func clientGo(id int){
    defer wg.Done()
    conn,err:=net.DialTimeout(SERVER_NETWORK,SERVER_ADDRESS,200*time.Millisecond)
    if err!=nil {
        printLog("connect server error ",err.Error())
        return
    }
    defer conn.Close()
    printLog("client",id ," connect to server. server addr is ",conn.RemoteAddr(),", local addr is ",conn.LocalAddr())


    requestNumber := 5
    conn.SetDeadline(time.Now().Add(5*time.Millisecond))
    for i := 0; i<requestNumber; i++ {
        _,err= writeTo(conn,fmt.Sprintf("%d",rand.Int31()))
        if err!=nil {
            printLog("Client write error ",err.Error())
            continue
        }

        printLog("send request to server")
    }
    reader:= bufio.NewReader(conn)
    for i := 0; i<requestNumber; i++  {
        data,err:=reader.ReadBytes(DELIMITER)
        if err!=nil {
            if err==io.EOF {
                printLog("the connect is closed")
            }else {
                printLog("Client read error ",err)
            }
            break
        }

        printLog("the server response is ",string(data))
    }
}