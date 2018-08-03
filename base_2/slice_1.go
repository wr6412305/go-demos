package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	// 一个切片在未初始化之前默认为nil，长度为0
	var num1 []int
	printSlice(num1)
	fmt.Println()

	// 切片的截取
	num2 := [] int {0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(num2)
	fmt.Println("num[1:4] =", num2[1:4])
	fmt.Println("num[:3] =", num2[:3])
	fmt.Println("num[4:] =", num2[4:])

	num3 := make([]int, 0, 5)
	printSlice(num3)

	num4 := num2[:2]
	printSlice(num4)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
