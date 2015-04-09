package pipes

import (
    "os/exec"
    "errors"
    "fmt"
    "bytes"
    "io"
)

func RunCmds(cmds []*exec.Cmd) ([]string,error){
    if cmds == nil || len(cmds)==0 {
        return nil,errors.New("The cmd slice is invalid!")
    }

    first := true
    var output []byte
    var err error
    for _,cmd := range cmds{
        fmt.Printf("Run command: %v ...\n",getCmdPlaintext(cmd))
        if !first {

            var stdinBuf bytes.Buffer
            stdinBuf.Write(output)
            cmd.Stdin = &stdinBuf
        }

        var stdoutBuf bytes.Buffer
        cmd.Stdout = &stdoutBuf
        if err= cmd.Start(); err!=nil {
            return nil,getError(err,cmd)
        }

        if err = cmd.Wait(); err != nil {
            return nil, getError(err, cmd)
        }

        output = stdoutBuf.Bytes()
        if first {
            first = false
        }
    }

    lines:= make([]string,0)
    var outputBuf bytes.Buffer
    outputBuf.Write(output)
    for{
        line,err:= outputBuf.ReadBytes('\n')
        if err!=nil {
            if err == io.EOF {
                break
            }else {
                return nil,getError(err,nil)
            }
        }
        lines = append(lines,string(line))
    }
    return lines,nil
}

func getCmdPlaintext(cmd *exec.Cmd) string{
    var buf bytes.Buffer
    buf.WriteString(cmd.Path)
    for _,arg := range cmd.Args[1:]{
        buf.WriteRune(' ')
        buf.WriteString(arg)
    }
    return buf.String()
}



func getError(err error, cmd *exec.Cmd, extraInfo ...string) error {
    var errMsg string
    if cmd != nil {
        errMsg = fmt.Sprintf("%s  [%s %v]", err, (*cmd).Path, (*cmd).Args)
    } else {
        errMsg = fmt.Sprintf("%s", err)
    }
    if len(extraInfo) > 0 {
        errMsg = fmt.Sprintf("%s (%v)", errMsg, extraInfo)
    }
    return errors.New(errMsg)
}