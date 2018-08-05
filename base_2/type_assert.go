package main

import "fmt"

// 类型断言用于提取接口的底层值

func assert(i interface{}) {
	// i的具体类型为int，使用i.(int)来提取底层的int值,如果i的具体类型不是int
	// 则会报错 panic: interface conversion: interface {} is string, not int
	s := i.(int)
	fmt.Println(s)
}

func assert1(i interface{}) {
	// 当i的底层类型不是int是，ok为false，s为0，程序不会报错
	s, ok := i.(int)
	fmt.Println(s, ok)
}

func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "hello"
	assert1(s)
	assert1(i)
}
