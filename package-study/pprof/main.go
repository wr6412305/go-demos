package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"go-demos/package-study/pprof/data"
)

func main() {
	go func() {
		log.Println(data.Add("https://github.com/liangjisheng"))
	}()

	http.ListenAndServe("127.0.0.1:8080", nil)
}
