package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area() of Circle(c1) =", c1.getArea())
}

// 该method属于Circle类型对象中的方法
func (c Circle) getArea() float64 {
	// c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
