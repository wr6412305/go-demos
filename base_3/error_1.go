package main

// type error interface {
//    Error() string
//}
// fmt.Println()在打印错误时，会在内部调用Error()string 方法来得到该错误的描述

import (
	"os"
	"fmt"
	"net"
)

func test_error_1() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

func test_error_2() {
	f, err := os.Open("/test.txt")
	// 断言底层结构体类型，使用结构体字段获得更多信息
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// 断言底层结构体类型，调用方法获取更多信息
func test_error_3() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

func main() {
	test_error_1()
	test_error_2()
	test_error_3()
}
