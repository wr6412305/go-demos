package main

// Go语言支持头等函数(First Class Function)的编程语言，可以把函数赋值给变量，也可以把
// 函数作为其他函数的参数或者返回值。

import "fmt"

func main() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a)

	// 直接调用匿名函数
	func() {
		fmt.Println("hello world first class function")
	}()

	// 向匿名函数传递参数
	func(str string) {
		fmt.Println("welcome", str)
	}("Gophers")
}
