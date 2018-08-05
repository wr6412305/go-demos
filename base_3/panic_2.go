package main

import (
	"fmt"
	"runtime/debug"
)

// 恢复后获得堆栈跟踪

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()		// 打印堆栈跟踪
	}
}

func a() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])	// panic,数组越界
	fmt.Println("normally returned from a")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}
