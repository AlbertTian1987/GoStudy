package main
import (
    "os"
    "syscall"
    "fmt"
    "os/signal"
    "sync"
    "time"
    "runtime/debug"
    "os/exec"
    "GoStudy/multiproc/pipes"
    "strconv"
    "strings"
)

func main() {
    go func() {
        time.Sleep(5*time.Second)
        sigSendingDemo()
    }()

    sigHandleDemo()
}

func sigHandleDemo() {

    sigRecv1 := make(chan os.Signal,1)
    sigs1:= []os.Signal{syscall.SIGINT,syscall.SIGQUIT}
    fmt.Printf("Set notification for %s... [sigRecv1]\n", sigs1)

    signal.Notify(sigRecv1,sigs1...)

    sigRecv2 :=make(chan os.Signal,1)
    sigs2 := []os.Signal{syscall.SIGQUIT}
    fmt.Printf("Set notification for %s... [sigRecv2]\n", sigs2)
    signal.Notify(sigRecv2,sigs2...)

    var wg sync.WaitGroup
    wg.Add(2)
    go func(){
        for sig:= range sigRecv1{
            fmt.Printf("Received a signal from sigRecv1: %s\n", sig)
        }
        fmt.Printf("End. [sigRecv1]\n")
        wg.Done()
    }()

    go func() {
        for sig := range sigRecv2 {
            fmt.Printf("Received a signal from sigRecv2: %s\n", sig)
        }
        fmt.Printf("End. [sigRecv2]\n")
        wg.Done()
    }()

    fmt.Println("Wait for 2 seconds... ")
    time.Sleep(2 * time.Second)
    fmt.Printf("Stop notification...")
    signal.Stop(sigRecv1)
    close(sigRecv1)
    fmt.Printf("done. [sigRecv1]\n")
    wg.Wait()

}

func sigSendingDemo() {
    defer func() {
        if err := recover();err!=nil {
            fmt.Printf("Fatal Error: %s\n",err)
            debug.PrintStack()
        }
    }()

    cmds := []*exec.Cmd{
        exec.Command("ps", "aux"),
        exec.Command("grep", "mysignal"),
        exec.Command("grep", "-v", "grep"),
        exec.Command("grep", "-v", "go run"),
        exec.Command("awk", "{print $2}"),
    }

    output,err:=pipes.RunCmds(cmds)
    if err != nil {
        fmt.Printf("Command Execution Error: %s\n", err)
        return
    }
    for _,pidS:=range output{
        pid,err:=strconv.Atoi(strings.TrimSpace(pidS))
        if err!=nil {
            continue
        }
        proc,err:=os.FindProcess(pid)
        if err!=nil {
            continue
        }
        sig := syscall.SIGQUIT
        fmt.Printf("Send signal '%s' to the process (pid=%d)...\n", sig, pid)
        err = proc.Signal(sig)
        if err != nil {
            fmt.Printf("Signal Sending Error: %s\n", err)
            return
        }
    }

}