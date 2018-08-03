package main

import "fmt"

// slice 可以看成是一种轻量级的数据结构，可以用来访问数组
// 的部分或者全部元素，而这个数组称之为slice的底层数组。
// slice有三个属性：指针，长度和容量。指针指向数组的第一个
// 可以从 slice 中访问的元素，这个元素不一定是数组的第一个
// 元素.长度指的是slice中的元素个数,不能超过slice的容量
// Go语言的内置函数len和cap用来返回slice的长度和容量

func print_info(my_slice[] int){
	fmt.Println("len:", len(my_slice))
	fmt.Println("cap:", cap(my_slice))
	for i, v := range my_slice{
		fmt.Println("element[", i, "] = ", v)
	}
}

func main(){
	my_slice01 := []int{1, 2, 3, 4, 5}
	my_slice02 := make([]int, 5)
	my_slice03 := make([]int, 5, 6)
	my_slice04 := append(my_slice03, 8, 9, 10)
	print_info(my_slice01)
	print_info(my_slice02)
	print_info(my_slice03)
	print_info(my_slice04)
	fmt.Println()

	// 结合Go的不定参数，使用append实现两个slice的拼接
	my_slice05 := append(my_slice01, my_slice04...)
	print_info(my_slice05)
}
