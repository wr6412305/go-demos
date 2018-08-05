package main

import "fmt"

// new会根据类型的大小，分配该大小的零值空间，然后返回该空间首地址
// make被编译器翻译成具体的创建函数，由其分配内存并初始化成员变量，返回对象，而非指针

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//println(arr)
	fmt.Println(arr)

	var slice1 = make([]int, 5, 10)
	//println(slice1)
	fmt.Println(slice1)

	slice2 := new([]int)	// 返回指针
	(*slice2) = append((*slice2), 1)
	fmt.Println(*slice2)
	fmt.Println(slice2)
	fmt.Println()

	var map1 = make(map[string]int, 1000)	// 分配内存并初始化
	fmt.Println(map1)

	map2 := new(map[string]int)		// 分配内存
	(*map2) = map[string]int{}		// 初始化
	(*map2)["hello"] = 4
	fmt.Println((*map2))
}
