package web
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
    if err!=nil {
        tt.Fatal(err)
    }else {
        err =t.Execute(os.Stdout,user)
        if err!=nil {
            tt.Fatal(err)
        }
    }
}



func Test_FuncMap2(tt *testing.T){
    t,err := template.ParseFiles("my.html","header.html","content.html","footer.html")
    if err!=nil {
        tt.Fatal(err)
    }else {
        t.ExecuteTemplate(os.Stdout,"my",nil)
        t.ExecuteTemplate(os.Stdout,"header",nil)
        t.ExecuteTemplate(os.Stdout,"content",nil)
        t.ExecuteTemplate(os.Stdout,"footer",nil)
    }
}