package main

import (
	"net/http"

	"web-demo/restful/routes"
)

func main() {
	r := routes.NewRouter()

	http.ListenAndServe(":9090", r)
}
