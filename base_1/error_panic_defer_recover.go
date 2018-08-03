package main

import "fmt"

// Go语言的错误处理流程:当一个函数在执行过程中出现了异常或遇到panic(),
// 正常语句就会立即终止,然后执行defer语句,再报告异常信息,最后退出goroutine
// 如果在defer中使用了recover()函数,则会捕获错误信息,使该错误信息终止报告
// 这里结合前面所学的error内容,实现一个自定义的Error()方法,并结合panic和
// recover,集中展示下Go语言的错误处理机制:

// 自定义错误类型
type ArithmeticError struct {
	error
}

// 重写Error()方法
func (this *ArithmeticError) Error() string {
	return "自定义的error, error名称为算法不合法"
}

// 定义除法运算
func Devide(num1, num2 int) int {
	if num2 == 0 {
		// //当然也可以使用ArithmeticError{}同时recover等到ArithmeticError类型
		panic(&ArithmeticError{})
	}else {
		return num1 / num2
	}
}

func main() {
	var a, b int
	fmt.Printf("input two number: ")
	fmt.Scanf("%d %d", &a, &b)

	defer func () {
		if r := recover(); r != nil {
			// %v表示打印结构体
			fmt.Printf("panic的内容%v\n", r)
			if _, ok := r.(error); ok {
				fmt.Println("panic--recover()得到的是error类型")
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
	fmt.Println("结果是:", rs)
}
