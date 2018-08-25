package main

import (
	"fmt"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, Welcome to go web programming...")
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}
