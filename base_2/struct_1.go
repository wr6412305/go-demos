package main

import "fmt"

// 虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型
// 比如在下面的Person结构体里，虽说字段是匿名的，但 Go 默认这
// 些字段名是它们各自的类型.所以Person结构体有两个名为string和int的字段
type Person struct {
	string
	int
}

func main() {
	var p1 Person
	p1.string = "Mike"
	p1.int = 50
	fmt.Println(p1)
}
