package main

import (
	"fmt"
)

func recover1() {
	defer func() {
		// 捕捉异常
		if p := recover(); p != nil {
			fmt.Printf("error: %v\n", p)
		}
		return
	}()

	if x := 0; x == 0 {
		// 异常捕捉后函数直接返回，而不是继续执行
		panic("x can not be 0")
	}

	fmt.Println("After error")
}
