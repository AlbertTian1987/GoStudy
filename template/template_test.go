package template
import (
    "testing"
    "text/template"
    "os"
)


func Test_FuncMap(tt *testing.T){
    f1 := Friend{"albert"}
    f2 := Friend{"John"}

    user := User{"田国辉",[]string{"e1@163.com","e2@163.com","e3@163.com"},[]Friend{f1,f2}}

    t:=template.New("t1.html")
    t = t.Funcs(template.FuncMap{"emailDeal":EmailDealWith})
    t,err := t.ParseFiles("t1.html")
    f,_:=os.Create("test.html")
    defer  f.Close()
    if err!=nil {
        tt.Fatal(err)
    }else {
        err =t.Execute(f,user)
        if err!=nil {
            tt.Fatal(err)
        }
    }
}