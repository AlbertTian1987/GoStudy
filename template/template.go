package template
import (
    "net/http"
    "text/template"
    "fmt"
    "strings"
)
type User struct {
    UserName string
    Emails []string
    Friends []Friend
}

type Friend struct {
    Name string
}

func Template1(rw http.ResponseWriter,req *http.Request){
    user := new(User)
    user.UserName = "田国辉"
    user.Emails = make([]string,4)
    user.Emails[0]="e1@163.com"
    user.Emails[1]="e2@163.com"
    user.Emails[2]="e3@163.com"
    user.Emails[3]="<script>alert('hehe')</script>"

    user.Friends = make([]Friend,3)
    user.Friends[0] = Friend{"约翰"}
    user.Friends[1] = Friend{"大卫"}
    user.Friends[2] = Friend{"安德森"}



    t,err:=template.New("base.html").Funcs(funcMap).ParseFiles("template/t1.html")
    if err!=nil {
        fmt.Println(err)
        rw.Write([]byte("hehe"))
    }else {
        t.Execute(rw,user)
    }
}

var funcMap = template.FuncMap {"EmailDeal":EmailDealWith}

func EmailDealWith(args ...interface{}) string{
    ok:=false
    var s string
    if len(args) ==1 {
        s,ok = args[0].(string)
    }
    if !ok {
        s = fmt.Sprint(args...)
    }

    substrs := strings.Split(s,"@")
    if len(substrs) !=2 {
        return s
    }

    return (substrs[0])+" at "+substrs[1]
}