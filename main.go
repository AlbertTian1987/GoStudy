package main
import (
    "net/http"
    "GoStudy/web"
    "GoStudy/db"
)


func main(){
    http.HandleFunc("/",web.Template1)
    http.HandleFunc("/rwCookie",web.SetAndGetCookie)
    http.ListenAndServe(":9090",nil)

    db.CloseDB()
}