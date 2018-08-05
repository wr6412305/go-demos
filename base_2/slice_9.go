package main

import (
	"fmt"
	"unsafe"
)

// 每次cap的改变，指向array的ptr就会变化一次
// 实际go在append的时候放大cap是有规律的。在 cap 小于1024的情况下是每次扩大到
// 2 * cap ，当大于1024之后就每次扩大到 1.25 * cap
// 所以上面的测试中cap变化是 1, 2, 4, 8

func main() {
	s := make([]int, 1)
	fmt.Printf("len: %d cap: %d array ptr: %v\n",
		len(s), cap(s), unsafe.Pointer(&s))
	fmt.Printf("len: %d cap: %d array ptr: %v\n",
		len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))
	fmt.Println()

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("len: %d cap: %d array ptr: %v\n",
			len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))
	}

	fmt.Println("Array:", s)
	fmt.Println()

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len: %d cap: %d array ptr: %v\n",
		len(s1), cap(s1), *(*unsafe.Pointer)(unsafe.Pointer(&s1)))
	fmt.Println("Array:", s1)

	s2 := s1[1:3]
	fmt.Printf("len: %d cap: %d array ptr: %v\n",
		len(s2), cap(s2), *(*unsafe.Pointer)(unsafe.Pointer(&s2)))
	fmt.Println("Array:", s2)
}
