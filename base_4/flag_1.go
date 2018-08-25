package main

// 命令行参数
// ./flag_1.exe -name=liangjisheng

import (
	"flag"
	"fmt"
)

func main() {
	// 第一个位参数名称，第二个位默认值，第三个是说明，现在定义了一个
	// 命令行标志name
	username := flag.String("name", "", "Input your username")
	// 调用flag.Parse()解析命令行参数到定义的flag
	flag.Parse()
	fmt.Println("Hello", *username)
}
