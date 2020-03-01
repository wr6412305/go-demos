package main

// curl "http://127.0.0.1:8080/sum?a=1&b=1"

import (
	"fmt"
	"go-demos/go-kit-demo/v1/v1endpoint"
	"go-demos/go-kit-demo/v1/v1service"
	"go-demos/go-kit-demo/v1/v1transport"
	"net/http"
)

func main() {
	server := v1service.NewService()
	endpoints := v1endpoint.NewEndPointServer(server)
	httpHandler := v1transport.NewHTTPHandler(endpoints)
	fmt.Println("server run 127.0.0.1:8080")
	_ = http.ListenAndServe("127.0.0.1:8080", httpHandler)
}
