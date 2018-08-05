package main

import "fmt"

// 上篇文章说了，所有类型的函数参数都是按值传递的
// 但Go中其实也是有传引用的地方，但不是函数的参数，而是闭包对外部环境是通过引用访问的

func main() {
	a := new(int)
	fmt.Println(a)
	func(){
		a = nil
	}()

	// 闭包中使用引用修改变量a
	fmt.Println(a)		// <nil>

	// output: 4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	defer fmt.Println()

	// 5 5 5 5 5
	// 通过闭包引用的方式输出变量i
	for i := 0; i < 5; i++ {
		defer func(){ fmt.Printf("%d ", i) } ()
	}
	defer fmt.Println()

	// 回避上面副作用的方法是通过参数传值或每次闭包构造不同的临时变量
	for i := 0; i < 5; i++ {
		i := i		// 每次构造不同的临时变量
		defer func() { fmt.Printf("%d ", i) }()
	}
	defer fmt.Println()

	// 通过参数传递
	for i := 0; i < 5; i++ {
		defer func (i int) { fmt.Printf("%d ", i) } (i)
	}

	defer fmt.Println()
}
