package main

import (
	"fmt"
)

// 自定义错误类型
type ArithmeticError struct {
	error
}

// 重写Error()方法
func (this *ArithmeticError) Error() string {
	return "self define error, param is unlegal."
}

func Devide(num1, num2 int) int {
	if 0 == num2 {
		panic(&ArithmeticError{})
	} else {
		return num1 / num2
	}
}

func err1() {
	var a, b int
	fmt.Printf("input two number:")
	fmt.Scanf("%d %d", &a, &b)

	// 输出10 表明是一个回车符，即\r,必须处理这个字符，否则下次在次输入的时候会默认读取\r
	var line rune
	fmt.Scanf("%c", &line)
	fmt.Println("line:", line) // 10

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic is %v\n", r)
			if _, ok := r.(error); ok {
				fmt.Println("panic--recover()得到的是error类型.")
			}
			if _, ok := r.(*ArithmeticError); ok {
				fmt.Println("panic-recover()得到的是ArithmeticError类型")
			}
			if _, ok := r.(string); ok {
				fmt.Println("panic-recover()得到的是string类型")
			}
		}
	}()

	rs := Devide(a, b)
	fmt.Println("结果是: ", rs)
}
