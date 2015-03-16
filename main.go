package main

import (
	"GoStudy/web"
	"net/http"
    _"net/http/pprof"
)

func main() {
    for i:=0;i<10000;i++{
        go web.F()
    }

	http.HandleFunc("/", web.Template1)
	http.HandleFunc("/pprof", web.Pprof)
	http.HandleFunc("/rwCookie", web.SetAndGetCookie)
	http.ListenAndServe(":9090", nil)
}
