package web

import (
	"net/http"
	"runtime/pprof"
)

var quit chan struct{} = make(chan struct{})

func F() {
	<-quit
}

func Pprof(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}
