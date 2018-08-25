package main

// 将字符串解析为整数，ParseInt 支持正负号，ParseUint 不支持正负号。
// base 表示进位制（2 到 36），如果 base 为 0，则根据字符串前缀判断，
// 前缀 0x 表示 16 进制，前缀 0 表示 8 进制，否则是 10 进制。
// bitSize 表示结果的位宽（包括符号位），0 表示最大位宽。
// func ParseInt(s string, base int, bitSize int) (i int64, err error)
// func ParseUint(s string, base int, bitSize int) (uint64, error)

import (
	"fmt"
	"strconv"
)

func main() {
	// 255 <nil>
	fmt.Println(strconv.ParseInt("FF", 16, 0))
	// 0 strconv.ParseInt: parsing "0xFF": invalid syntax
	fmt.Println(strconv.ParseInt("0xFF", 16, 0))
	// 255 <nil>
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))
	fmt.Println(strconv.ParseInt("9", 10, 4))
	fmt.Println(strconv.ParseInt("9", 10, 0))
}
