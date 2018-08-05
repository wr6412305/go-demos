package main

import "fmt"

// 引用类型和指针在底层实现上是一样的.
// 但是引用类型在语法上隐藏了显示的指针操作
// 可以自己给Go语言模拟一个引用类型，可以将值类型特定的数组
// 定义为一个引用类型(同时提供一个构造函数)

type RefIntArray2 *[2]int

func NewRefIntArray2() RefIntArray2{
	return RefIntArray2(new([2]int))
}

// 可以将RefIntArray2当做引用类型来使用

func main() {
	refarr2 := NewRefIntArray2()
	fmt.Println(refarr2)
	modify(refarr2)
	fmt.Println(refarr2)
}

func modify(arr RefIntArray2) {
	arr[0] = 1
}
