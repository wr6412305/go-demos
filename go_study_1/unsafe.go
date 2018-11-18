package main

import (
	"fmt"
	"unsafe"
)

func unsafe1() {
	x := 42
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(&x))
	fmt.Println(unsafe.Sizeof(
		struct {
			x int
			y float32
		}{}))
}

// unsafe.Alignof()接收一个可产生任何类型的表达式（表达式不参与求值）
// 返回使这个类型对齐时所需要的字节数
// unsafe.OffsetOf接收一个结构体字段选择器，返回该字段相对于结构体开始地址的偏移量
func unsafe2() {
	var x struct {
		a bool
		b int16
		c []int
	}

	fmt.Printf("Sizeof(x)\t= %d\tAlignof(x)\t=%d\n",
		unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Printf("Sizeof(x.a)\t= %d\tAlignof(x.a)\t= %d\tOffsetof(x.a)\t= %d\n",
		unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Printf("Sizeof(x.b)\t= %d\tAlignof(x.b)\t= %d\tOffsetof(x.b)\t= %d\n",
		unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Printf("Sizeof(x.c)\t= %d\tAlignof(x.c)\t= %d\tOffsetof(x.c)\t= %d\n",
		unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))
}
