package main

import "fmt"

// 当函数只有一个值参数，它只能接受一个值参数
// 当一个方法有一个值接收器，它可以接受值接收器和指针接收器

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width: 5,
	}
	area(r)
	r.area()

	p := &r
	// area(p)		// compilation error
	// 这时编译器会把p.area()解释为(*p).area()
	p.area()		// 通过指针调用值接收器
}
