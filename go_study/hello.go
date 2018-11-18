package main

import "fmt"

func hello() {
	// 左大括号必须和if在同一行，而且大括号不能省略
	if 0 == 7%2 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 只有if条件的情况
	if 0 == 8%4 {
		fmt.Println("8 is divisible by 4")
	}

	// if条件可以包含一个初始化表达式，这个表达式中的变量
	// 是这个条件判断结构的局部变量
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
