package main

import (
	"fmt"
)

// 字符串实际上就是类型为byte的只读切片

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println()

	fmt.Println("Printf with %x:")
	fmt.Printf("% x\n", sample)

	// 使用q标识，可以将字符串中任何不可打印的字节序列转义
	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}
