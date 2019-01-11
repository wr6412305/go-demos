package main

import (
	"net/http"
	"web-demo/microservice/movies/routes"
)

func main() {
	r := routes.NewRouter()
	http.ListenAndServe(":8001", r)
}
