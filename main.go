package main
import (
    "net/http"
    "GoStudy/template"
)


func main(){
    http.HandleFunc("/",template.Template1)
    http.ListenAndServe(":9090",nil)
}