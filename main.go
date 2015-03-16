package main

import (
	"GoStudy/db"
	"GoStudy/web"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.Template1)
	http.HandleFunc("/rwCookie", web.SetAndGetCookie)
	http.ListenAndServe(":9090", nil)

	db.CloseDB()
}
