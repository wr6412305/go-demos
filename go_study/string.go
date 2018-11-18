package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\n")
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Printf("\n")
}

func printRunes(s string) {
	// rune表示一个Unicode码点,无论一个码点被编码为多少个字节，都可以表示为一个rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Printf("\n")
}

func printCharsAndBytes(s string) {
	// range返回一个rune的位置以及rune本身
	for index, r := range s {
		fmt.Printf("%c starts at byte %d\n", r, index)
	}
}

func string1() {
	name := "Hello World"
	printBytes(name)
	printChars(name)
	printCharsAndBytes(name)

	name = "Señor"
	printBytes(name)
	printChars(name)
	printRunes(name)
	printCharsAndBytes(name)
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func string2() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str) // Café
	length(str)

	byteSlice = []byte{67, 97, 102, 195, 169}
	str = string(byteSlice)
	fmt.Println(str) // Café

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str) // Señor
	length(str)
}

// 字符串是不可变的，如果要改变字符串的内容，需要转换为rune切片，修改切片中的内容，
// 再转换为string返回
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func string3() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
