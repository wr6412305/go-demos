package main

import (
	"net/http"
	"web-demo/microservice/users/routes"
)

func main() {
	r := routes.NewRouter()
	http.ListenAndServe(":8000", r)
}
