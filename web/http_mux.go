package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, myroute!")
}

// 实现一个简易的路由器
func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
