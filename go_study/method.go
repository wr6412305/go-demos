package main

import (
	"fmt"
	"math"
)

type employee struct {
	name     string
	salary   int
	currency string
}

func (e employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func method1() {
	emp1 := employee{
		name:     "ljs",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary()
}

// Go使用方法的2点原因
// 1.Go 不是一个纯面向对象的编程语言，它不支持class类型.因此通过在一个类型上
// 建立方法来实现与class相似的行为
// 2.同名方法可以定义在不同类型上，但Go不允许同名函数(没有函数重载)

type Rectangle struct {
	length int
	width  int
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

func method2() {
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())
}

// 函数有一个值参数时，它只接受一个值参数
// 方法有一个值接受者时，它可以接受值和指针接受者
type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area function result: %d\n", r.length*r.width)
}

func (r rectangle) area() {
	fmt.Printf("Area method result: %d\n", r.length*r.width)
}

func method3() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	// area(p)	// error: param type error
	p.area() // Go会默认转换为(*p).area()
}

// 与值参数相似，一个接收指针参数的函数指针接收指针， 而一个接收者为指针的方法可以
// 接受值接收者和指针接收者
