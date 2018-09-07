package main

import (
	"fmt"
	"strconv"
)

// 任何实现了String方法的类型都能作为参数被fmt.Println调用

type Human struct {
	name string
	age int
	phone string
}

// Human类型实现了fmt.Stringer接口
func (h Human) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
