package web
import (
    "net/http"
    "fmt"
    "time"
)


func SetAndGetCookie(rw http.ResponseWriter,req *http.Request){

    if c_name,err:=req.Cookie("username");err == nil {
        fmt.Fprintln(rw,c_name)
    }else{
        expiration:=time.Date(2015,7,13,0,0,0,0,time.Local)
        c := http.Cookie{Name:"username",Value:"Albert",Expires:expiration}
        http.SetCookie(rw,&c)
    }
}
