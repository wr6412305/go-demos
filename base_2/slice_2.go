package main

import "fmt"

// append and copy

func main() {
	var numbers []int
	printSlice(numbers)

	// 添加一个元素
	numbers = append(numbers, 0)
	printSlice(numbers)
	numbers = append(numbers, 1)
	printSlice(numbers)
	// 同时添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	// 创建切片number1是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), cap(numbers) * 2)
	// 拷贝numbers的内容到numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
