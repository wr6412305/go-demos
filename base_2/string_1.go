package main

import "fmt"

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		// 若s中包含像俄文、希腊文等字符时，他们有的字符UTF-8编码会占用2个字节
		// %c默认为一个字节，此时会有问题
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
}

// rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune
// 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示
// 让我们修改一下上面的程序，用 rune 来打印字符
func printChars1(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}

// for range循环遍历string
func printCharsAndBytes(s string) {
	// 循环返回的是是当前 rune 的字节位置
	for index, rune := range(s) {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
	fmt.Println()
}

func main() {
	name := "Hello World"
	printBytes(name)
	printChars(name)
	printChars1(name)

	name = "Señor"
	printChars(name)
	printChars1(name)

	// 可以看出，第三个字符占2个字节
	printCharsAndBytes(name)
}
