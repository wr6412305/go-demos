package main

import (
	"fmt"
)

// 指针方法和值方法都可以在指针或非指针上被调用，Go会自动转换

type Info struct {
	address string
	phone   string
}

// 接受者是非指针
func (this Info) changeInfo() {
	this.address = "上海"
	this.phone = "10086"
	fmt.Println(this)
}

// 接受者是指针
func (this *Info) changeInfo2() {
	this.address = "上海"
	this.phone = "10086"
	fmt.Println(this)
}

func demo() {
	info := &Info{"北京", "10086"}
	//指针调用，但接收者自动被转换成了非指针，changeInfo函数应用的是info的拷贝
	info.changeInfo()
	fmt.Println(info)

	info2 := &Info{"北京", "10086"}
	//非指针调用，但接收者自动被转换成了指针，changeInfo函数应用的是真正的info实例
	info2.changeInfo2()
	fmt.Println(info2)
}

func main() {
	demo()
}
