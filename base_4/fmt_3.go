package main

import (
	"fmt"
)

// 功能同 Sprintf，只不过结果字符串被包装成了 error 类型。
// func Errorf(format string, a ...interface{}) error

func main() {
	// fmt.Print会在非字符串参数之间会添加空格，返回写入的字节数
	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	// 所有参数之间添加空格
	fmt.Println("a", "b", 1, 2, 3, "c", "d")
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)

	if err := percent(30, 70, 90, 160); err != nil {
		fmt.Println(err)
	}
}

func percent(i ...int) error {
	for _, n := range i {
		if n > 100 {
			return fmt.Errorf("数值%d超出范围(100)", n)
		}
		fmt.Print(n, "%\n")
	}
	return nil
}
