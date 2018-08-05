package main

import (
	"fmt"
)

// 属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样

type address struct {
	city, state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

type Person struct {
	firstName, lastName string
	address		// 匿名字段
}

func main() {
	p := Person{
		firstName: "Elon",
		lastName: "Musk",
		address : address{
			city: "Los Angeles",
			state: "California",
		},
	}

	p.fullAddress()		// 访问address结构体的fullAddress方法
	p.address.fullAddress()		// 也可以明确的调用
}
