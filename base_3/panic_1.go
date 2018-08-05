package main

// 当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制
// 返回到该函数的调用方。这样的过程会一直持续下去，直到当前协程的所有函数
// 都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪，最后程序终止

import (
	"fmt"
	)

// panic有两个合理的用例
// 1.发生了一个不能回复的错误，此时程序不能继续运行
// 2.发生了一个编程上的错误

func fullname(firstname *string, lastname *string) {
	defer fmt.Println("deferred call in fullName")
	if firstname == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastname == nil {
		panic("runtime error: last name cannot be nil")
	}

	fmt.Printf("%s %s\n", *firstname, *lastname)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullname(&firstName, nil)
	fmt.Println("returned normally from main")
}
