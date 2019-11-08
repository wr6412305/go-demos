package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	cf := "server.crt"
	ck := "server.key"
	http.ListenAndServeTLS("127.0.0.1:8080", cf, ck, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	fmt.Fprintf(w, "%s", "https request")
}
