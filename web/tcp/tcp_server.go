package main
import (
    "net"
    "fmt"
)
func main(){
    Server()
}

func Server() {
    listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("127.0.0,1"), 11990, ""})
    if err!=nil {
        fmt.Println("监听端口失败:", err.Error())
        return
    }
    fmt.Println("已经初始化连接，等待客户端连接...")
    for {
        conn, err := listener.AcceptTCP()
        if err!=nil {
            fmt.Println("接受客户端连接异常：", err.Error())
            continue
        }

        fmt.Println("客户端连接来自：", conn.RemoteAddr().String())

        go func(conn *net.TCPConn) {
            data := make([]byte, 1024)
            for {
                i,err := conn.Read(data)
                fmt.Println("客户端发来数据：",string(data[:i]))
                if err!=nil {
                    fmt.Println("读取客户端数据错误:",err.Error())
                    break
                }

                conn.Write([]byte{'f','i','n','i','s','h'})
            }
        }(conn)
    }

}