package main
import (
    "fmt"
    "os/exec"
    "bufio"
    "bytes"
    "io"
    "GoStudy/multiproc/pipes"
)

func main() {
    demo1()
    fmt.Println()
    demo2()

    demo3()
}
func demo1() {
    useBufferIo := false
    fmt.Println("Run command `echo -n \"My first command from golang.\"`: ")
    cmd0 := exec.Command("echo", "-n", "My first command from golang.")
    stdout0, err := cmd0.StdoutPipe()
    if err != nil {
        fmt.Printf("Error: Can not obtain the stdout pipe for command No.0: %s\n", err)
        return
    }
    if err := cmd0.Start(); err != nil {
        fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
        return
    }
    if !useBufferIo {
        var outputBuf0 bytes.Buffer
        for {
            tempOutput := make([]byte, 5)
            n, err := stdout0.Read(tempOutput)
            if err != nil {
                if err == io.EOF {
                    break
                } else {
                    fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
                    return
                }
            }
            if n > 0 {
                outputBuf0.Write(tempOutput[:n])
            }
        }
        fmt.Printf("%s\n", outputBuf0.String())
    } else {
        outputBuf0 := bufio.NewReader(stdout0)
        output0, _, err := outputBuf0.ReadLine()
        if err != nil {
            fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
            return
        }
        fmt.Printf("%s\n", string(output0))
    }
}
func demo2(){

    var stdoutBuf bytes.Buffer
    var stdinBuf bytes.Buffer
    var data []byte

    cmd1 := exec.Command("ls", "-al")
    cmd1.Stdout = &stdoutBuf

    cmd1.Start()
    cmd1.Wait()

    data = stdoutBuf.Bytes()
    fmt.Println("ls ---------")
    fmt.Println(string(data))

    stdoutBuf.Reset()

    cmd2 := exec.Command("grep","apipe")

    stdinBuf.Write(data)
    cmd2.Stdin = &stdinBuf
    cmd2.Stdout = &stdoutBuf
    cmd2.Start()
    cmd2.Wait()

    data = stdoutBuf.Bytes()
    fmt.Println("grep ---------")
    fmt.Println(string(data))


}

func demo3() {
        cmds := []*exec.Cmd{
            exec.Command("ps","aux"),
            exec.Command("grep","apipe"),
        }
        strings,err:=pipes.RunCmds(cmds)
        if err!=nil {
            fmt.Println(err)
            return
        }

        for _,s:=range strings{
            fmt.Print(s)
        }
}
