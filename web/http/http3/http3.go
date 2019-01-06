package main

import "net/http"

type Person struct {
	pName string
}

func (p *Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("person name: " + p.pName))
}

func main() {
	personOne := &Person{pName: "ljs"}

	// 任何路由进来好像都会调用Person的ServeHTTP的方法
	http.ListenAndServe(":5500", personOne)
}
