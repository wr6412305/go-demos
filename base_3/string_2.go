package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// 自定义函数实现shell中的basename函数
// fmt.Println(basename("a/b/c.go")) // "c"
// fmt.Println(basename("c.d.go"))   // "c.d"
// fmt.Println(basename("abc"))      // "abc"

func basename(s string) string {
	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserver everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename_1(s string) string {
	// 使用strings.LastIndex
	slash := strings.LastIndex(s, "/") // -1 if not founc
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// 这个例子的目标是获取一个整数的字符串表现形式，例如"12345"
// 然后每隔3个字符插入一个逗号，"12,345"
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

// bytes包提供了Buffer类型来高效的操作字节切片。Buffer
// 最开始是空的，之后会随着string,byte,[]byte类型数据的
// 写入逐步增长.bytes.Buffer变量是不需要初始化的，因为
// 零值一样可用
func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func test_strconv() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Println(y, strconv.Itoa(x))
	// 使用不同的基数来格式化字符串
	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)
	s = fmt.Sprintf("x=%d", x)
	fmt.Println(s)
	s = fmt.Sprintf("x=%x", x)
	fmt.Println(s)

	n, err := strconv.Atoi("123")              // n is int
	n1, err := strconv.ParseInt("123", 10, 64) // 基数是10，最大是64bit(int64)
}

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename_1("a/b/c.go"))
	fmt.Println(comma("1234541654654131362"))
	fmt.Println(IntsToString([]int{1, 2, 3}))
	test_strconv()
}
