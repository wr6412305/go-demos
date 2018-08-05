package main

import "fmt"

// 没有包含方法的接口称为空接口。空接口表示为interface{}.由于空接口没有方法
// 因此所有类型都实现了空接口

// 使用空接口作为参数，因此可以传递任何类型给describe()函数
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "hello, world"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Mike",
	}
	describe(strt)
}
