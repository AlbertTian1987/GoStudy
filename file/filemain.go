package main
import (
    "os/exec"
    "os"
    "path/filepath"
    "fmt"
)

func main() {
    file, _ := exec.LookPath(os.Args[0])
    path:= filepath.Dir(file)
    fmt.Println(path)
}
