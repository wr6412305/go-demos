package main

import "fmt"

// Go语言支持匿名函数，可作为闭包.匿名函数是一个内联语句或表达式.优越性在于
// 可以直接使用函数内的变量,不必声明.

// getSequence这个函数没有参数,返回值为一个函数,这个函数没有参数,返回类型为int
func getSequence () func () int {
	i := 0
	return func () int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()

	// 调用nextNumber函数,i变量自增1并返回
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// 创建新的函数nextNumber1,并查看结果
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
