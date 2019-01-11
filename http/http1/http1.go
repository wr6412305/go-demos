package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/about", aboutFunc)

	// nil表示使用默认mux
	http.ListenAndServe(":15500", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello go web app."))
}

func aboutFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello about page."))
}
