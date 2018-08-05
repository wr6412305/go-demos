package main

import "fmt"

// 类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。
// 它与一般的 switch 语句类似。唯一的区别在于类型选择指定的是类型，
// 而一般的 switch 指定的是值

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("type is string")
	case int:
		fmt.Println("type is int")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	findType("hello")
	findType(77)
	findType(89.98)
}
