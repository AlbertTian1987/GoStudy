package main
import (
    "io"
    "fmt"
    "time"
)


func inMemorySyncPipe() {
    reader,writer:=io.Pipe()
    go func() {
        output := make([]byte,100)
        n,err:= reader.Read(output)
        if err!=nil {
            fmt.Printf("Error:Can not read data from the named pipe: %s\n",err)
            return
        }

        fmt.Printf("Read %d byte(s) . [in-memory pipe]\n",n)
    }()

    input:= make([]byte,26)
    for i:=65;i<=90 ;i++  {
        input[i-65]= byte(i)
    }
    n,err := writer.Write(input)
    if err!=nil {
        fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
        return
    }
    fmt.Printf("Written %d byte(s). [in-memory pipe]\n", n)
    time.Sleep(200 * time.Millisecond)
}


func main(){
    inMemorySyncPipe()
}
