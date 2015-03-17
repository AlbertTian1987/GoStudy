package main
import (
    "fmt"
    "net"
)


func main(){

    conn,err := net.Dial("tcp","127.0.0.1:11990")
    if err!=nil {
        fmt.Println("连接服务器失败",err.Error())
        return
    }

    fmt.Println("连接服务器成功")
    defer conn.Close()

    sms := make([]byte,1024)
    for{
        fmt.Println("请输入要发送的消息：")
        _,err := fmt.Scan(&sms)
        if err!=nil {
            fmt.Println("数据输入异常：",err.Error())
            continue
        }

        conn.Write(sms)
        buf := make([]byte,128)
        c,err:=conn.Read(buf)
        if err != nil {
            fmt.Println("读取服务器数据异常:", err.Error())
        }
        fmt.Println(string(buf[0:c]))
    }

}
