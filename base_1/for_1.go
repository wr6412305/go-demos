package main

import "fmt"

// go语言中只有一个for循环这一种循环结构

func main(){

	// 基于计数器的for循环
	sum := 0
	// {左大括号必须和for在同一行，而且大括号不能省略
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()

	// 基于条件判断的for循环,类似于while循环
	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Println("The variable i is now:", i)
	}
	fmt.Println()

	// for-range结构, 用range来遍历一个数组
	arr := [...]int{6, 7, 8}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println()

	// Go语言有几种无限循环
	// i := 0; ; i++
	// for {}
	// for ;; {}
	// for true {}
}
