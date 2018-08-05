package main

import "fmt"

// 函数使用指针只接受指针，而使用指针接收器的方法可以使用值接收器和指针接收器
type rectangle struct {
	length int
	width  int
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width: 5,
	}

	p := &r
	perimeter(p)
	p.perimeter()

	// perimeter(r)		// error
	// 编译器会解释为(&r).perimeter()
	r.perimeter()		// 使用值来调用指针接收器
}
