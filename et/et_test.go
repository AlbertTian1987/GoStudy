package et
import (
    "fmt"
    "testing"
)

func ExampleM(){
    fmt.Print("hello world")
    //Output: hello world
}

func ExampleM2(){
    for i := 0; i<3;i++  {
        fmt.Println("hello world",i)
    }
    //Output: hello world 0
    //hello world 1
    //hello world 2
}

func TestTypeCategoryof(t *testing.T){
    str:= TypeCategoryof(87)
    if str!="integer" {
        t.FailNow()
    }

    str = TypeCategoryof("helklo")
    if str!="string" {
        t.FailNow()
    }

    str  = TypeCategoryof(87.33213)
    if str!="float" {
        t.FailNow()
    }

    str = TypeCategoryof(complex(1,3))
    if str!="complex" {
        t.FailNow()
    }

    str = TypeCategoryof(false)
    if str!="boolean" {
        t.FailNow()
    }

    str = TypeCategoryof(make(map[string]string))
    if str!="unknown" {
        t.FailNow()
    }
}