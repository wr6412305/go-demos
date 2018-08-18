package main

// rune能处理一切的字符，而byte仅仅局限在ascii

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello 世界"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	// 转换成rune类型来处理unicode字符
	str2 := []rune(str)
	fmt.Println(len(str2))
	str2[0] = 'w'
	fmt.Println(str2)         // 会输出整数编码
	fmt.Println(string(str2)) // 转换为字符串

	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c ", str2[i])
	}
	fmt.Println()

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", (str[i]))
	}
	fmt.Println()

	// range会隐式的unicode解码
	for _, v := range str {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
}
