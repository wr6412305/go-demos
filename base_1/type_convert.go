package main

import "fmt"

func main(){
	var sum int = 17
	var count int = 5
	var mean float32

	// go 语言中没有隐式类型转换，只能使用强制类型转换
	mean = float32(sum) / float32(count)
	fmt.Println("mean的值为:", mean)
}
