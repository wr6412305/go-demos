package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println()

	s1 := "hello, 世界"
	fmt.Println(len(s))                     // 返回字节长度13
	fmt.Println(utf8.RuneCountInString(s1)) // 返回码点长度，也就是字符长度9

	// utf8解码，得到每一个码点
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// 当range的对象是一个字符串时，隐式的调用了UTF-8解码，生成一个一个Unicode码点(rune)
	for i, r := range s1 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println()

	// []rune可以把字符串转换为Unicode码点
	fmt.Printf("% x\n", s1)
	r := []rune(s1)
	fmt.Printf("%x\n", r)

	// 如果将[]rune转换为一个字符串，会生成一个UTF8编码的字符串
	fmt.Println(string(r))

	// 将一个整数转为字符串时会把整数解释成rune值(rune的底层类型是int32)，生成UTF-8形式
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	// 如果rune不合法，会用错误字符替代显示
	fmt.Println(string(1234567))
}
