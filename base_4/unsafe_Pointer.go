package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// 和 pb := &x.b 等价
	// 将x的地址加上字段b的偏移量将地址转换为unsafe.Pointer，
	// 再转换为*int16
	// uintptr类似于一个地址整数
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))

	// 通过指针更新x.b
	*pb = 42
	fmt.Println(x.b)
}
