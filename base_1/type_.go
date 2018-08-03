package main

import "fmt"

var c, d int = 1, 2
var e, f = 123, "hello"

func main(){
	var v21 int32			// v21初始为0
	var v22 int = 2
	var v23 = 3			// 编译器自动推导类型
	v24 := 4			// 简易声明&定义的方式
	v21 = int32(v23)	// 强制转换

	g, h := 123, "hello"
	fmt.Println("v21 is", v21)
	fmt.Println("v22 is", v22)
	fmt.Println("v23 is", v23)
	fmt.Println("v24 is", v24)
	fmt.Println(c, d, e, f, g, h)
}
