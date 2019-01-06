package main

import (
	"net/http"
)

func main() {
	// 可以自己创建一个路由服务来代替使用默认 mux 服务
	mMux := http.NewServeMux()
	mMux.HandleFunc("/", rootFunc)

	http.ListenAndServe(":5500", mMux)
}

func rootFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go web app."))
}
