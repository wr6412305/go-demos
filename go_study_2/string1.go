package main

import (
	"fmt"
)

func string1() {
	s := "我" // utf-8编码中文为3个字节
	fmt.Printf("s的类型是: %t, 长度为: %d, 16进制为: %x\n", s, len(s), s)

	for i, b := range []byte(s) {
		fmt.Printf("第%d个字节为: %b\n", i, b)
	}
}
