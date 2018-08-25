package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpus:", runtime.NumCPU())    // 返回逻辑cpu的个数
	fmt.Println("goroot:", runtime.GOROOT())  // 返回GOROOT路径
	fmt.Println("os/platform:", runtime.GOOS) // 目标操作系统
}
