package main
import (
    "fmt"
    "time"
)

func main() {
    unbuf := make(chan byte)
    go func() {
        fmt.Println("Sleep one second .")
        time.Sleep(1*time.Second)
        num:= <- unbuf
        fmt.Println(fmt.Sprintf("receive from unbuf channel ,num is %d",num))
    }()
    num := byte(127)
    fmt.Println("Send num",num)
    unbuf<-num
    fmt.Println("Done")

}
