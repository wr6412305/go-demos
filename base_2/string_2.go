package main

import (
	"fmt"
	"unicode/utf8"
)

func length(s string) {
	// RuneCountInString返回字符串中rune的数量
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

// 字符串是不可变类型，但可以将字符串转换为一个rune切片，做改变，然后在转换为一个字符串
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	// 用字节切片构造字符串,包含用UTF-8编码后的16进制字节
	bytesSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(bytesSlice)
	fmt.Println(str)

	// 把16进制换成10进制
	bytesSlice1 := []byte{67, 97, 102, 195, 169}
	str1 := string(bytesSlice1)
	fmt.Println(str1)

	// 用rune切片构造字符串
	// runeSlice 包含字符串 Señor的 16 进制的 Unicode 代码点
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str2 := string(runeSlice)
	fmt.Println(str2)

	length(str1)
	length(str2)

	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
