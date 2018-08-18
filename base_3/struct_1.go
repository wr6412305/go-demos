package main

import (
	"fmt"
)

type language []string

type human struct {
	name string
	age  int
	sec  bool
}

type programmer struct {
	human
	company string
	language
}

func (h *human) introduce() {
	fmt.Println("人类的")
}

func (p *programmer) introduce() {
	fmt.Println("程序员的")
}

func main() {
	p := programmer{human: human{name: "sda", age: 123, sec: true}, company: "iflytek"}
	p.introduce()
	p.human.introduce()
}
