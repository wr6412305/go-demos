package main

import (
	"net/http"
	"web-demo/microservice/showtimes/routes"
)

func main() {
	r := routes.NewRouter()
	http.ListenAndServe(":8002", r)
}
