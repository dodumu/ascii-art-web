package main

import (
	"ascii-web/web"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("web/style"))
	http.Handle("/web/style/", http.StripPrefix("/web/style/", fs))

	http.HandleFunc("/", web.RequestHandler)
	http.HandleFunc("/ascii-art", web.ResultHandler)

	http.ListenAndServe(":8080", nil)
}
