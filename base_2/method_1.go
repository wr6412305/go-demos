package main

import (
	"math"
	"fmt"
)

// 为什么要使用方法？
// 1.Go不是纯粹的面向对象编程语言，而且Go不支持类，因此基于类型的方法是一种实现和类相似行为的途径
// 2.相同名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的

type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle{
		length: 10,
		width: 5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius:12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())
}
