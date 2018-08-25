package main

import (
	"fmt"
	"unsafe"
)

// 在Golang中，[]int 和 []MyInt是两种不同的类型。 因此，
// []int的值不能转换为[]MyInt，反之亦然。 但是在
// unsafe.Pointer的帮助下，转换是可能的

func main() {
	type MyInt int
	a := []MyInt{0, 1, 2}
	// b := ([]int)(a) // error: cannot convert a (type [MyInt]) to type []int
	b := *(*[]int)(unsafe.Pointer(&a))

	b[0] = 3
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	a[2] = 9
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
