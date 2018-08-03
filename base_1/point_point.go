package main

import "fmt"

// 二级指针

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 10
	ptr = &a
	pptr = &ptr

	fmt.Printf("变量a = %d\n", a)
	fmt.Printf("指针变量*ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
