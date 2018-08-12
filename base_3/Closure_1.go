package main

import "fmt"

// 闭包是匿名函数的一个特例，当一个匿名函数所访问的变量定义在函数体的外部时，
// 称这样的匿名函数为闭包

func main() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
	fmt.Println()

	fa := appendStr()
	fb := appendStr()
	fmt.Println(fa("world"))
	fmt.Println(fb("everyone"))
	fmt.Println(fa("Gopher"))
	fmt.Println(fb("!"))
}

// 每一个闭包都会绑定一个它自己的外围变量（Surrounding Variable）
func appendStr() func (string) string {
	t := "hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
