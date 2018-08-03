package main

import (
	"fmt"
	"unsafe"
)

const (
	c1 = "abcabcabcabcabcabc"
	c2 = len(c1)
	// 字符串类型在go里是个结构, 包含指向底层数组的指针和长度
	// 这两部分每部分都是8个字节,所以字符串类型大小为16个字节
	c3 = unsafe.Sizeof(c1)
)

func main(){
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d\n", area)
	println()
	println(a, b, c)

	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值
	// 常量表达式中,函数必须是内置函数,否则编译不过:
	println(c1, c2, c3)
}
