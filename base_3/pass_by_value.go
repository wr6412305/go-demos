package main

// 所有类型的函数参数都是传值的，并不是传引用
// 包括slice/map/chan等内置类型和自定义类型都是传值的
// 但因为slice和map/chan底层结构的差异，导致他们传值的影响并不完全相同
// 按值传递的slice只能修改其指针指向的数据，其他都不能修改
// slice是结构体和指针的混合体

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)		// [1 2 3]
	modifySlice(a)
	// 两次输出结果一样，说明函数调用并没有改变a，说明slice是传值的
	// 但slice本身是引用类型
	fmt.Println(a)		// [1 2 3]

	modifySlice1(a)
	// modifySlice1(a)修改的a中的元素值，但slice还是传值的，slice中的指向数据的指针按值传递
	// 通过函数修改参数内容的机制有很多，也可以传地址，并不能通过这样就认为slice是传引用的
	fmt.Println(a)		// [0 2 3]
	fmt.Println()

	pint := new(int)
	fmt.Println(pint)
	modify(pint)
	// 两次输出结果一样，指针pint本身并没有什么变化，因为它是按值传递的，
	// 传指针或传地址只能修改指针指向的内存的值
	fmt.Println(pint)
}

func modifySlice(data []int) {
	data = nil
}

func modifySlice1(data []int) {
	data[0] = 0
}

func modify(a *int) {
	a = nil
}
