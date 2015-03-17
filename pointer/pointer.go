package main
import (
    "fmt"
    "unsafe"
)


type Person struct {
    Name string `json:"name"`
    Age byte `json:"age"`
    Address string `json:"address"`
}

func main() {
    pp := &Person{
        Name :"提昂",
        Age:18,
        Address:"海淀区dddxz",
    }

    fmt.Println(*pp)

    var ppUptr = uintptr(unsafe.Pointer(pp))
    var namePtr *string = (*string)(unsafe.Pointer(ppUptr+unsafe.Offsetof(pp.Name)))

    *namePtr="kity"

    var agePtr *byte= (*byte)(unsafe.Pointer(ppUptr+unsafe.Offsetof(pp.Age)))
    *agePtr = 245
    fmt.Println(*pp)

}
