package main
import (
    "net/http"
    "GoStudy/web"
)


func main(){
    http.HandleFunc("/",web.Template1)
    http.ListenAndServe(":9090",nil)
}