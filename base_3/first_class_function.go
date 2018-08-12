package main

import "fmt"

// 用户自定义函数类型
type add func(a int, b int) int

// 高阶函数:满足下列条件之一的函数
// 1.接收一个或多个函数做为参数
// 2.返回值是一个函数
func sample(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

// 返回一个函数
func sample_1() func(a, b int) int {
	f := func (a, b int) int {
		return a + b
	}
	return f
}

func main() {
	var a add = func(a int, b int) int {
		return a + b
	}

	s := a(5, 6)
	fmt.Println("sum", s)

	f := func(a, b int) int {
		return a + b
	}
	sample(f)

	f1 := sample_1()
	fmt.Println(f1(60, 5))
}
