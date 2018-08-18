package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	x int
	y int
}

func test_demo1() {
	intr1 := new(Point)
	intr1.x = 2
	intr1.y = 6
	intr2 := &Point{0, 3}
	intr3 := &Point{x: 5, y: 1}
	intr4 := &Point{y: 5}

	fmt.Println(intr1)
	fmt.Println(intr2)
	fmt.Println(intr3)
	fmt.Println(intr4)
}

// 使用工厂模式创建结构体实例
func NewPoint(x int, y int) *Point {
	if x < 0 || y < 0 {
		return nil
	}
	return &Point{x, y}
}

// 结构体字段后面的字符串表示字段的标签
// 反射获取结构体的标签
type Point1 struct {
	x int "this is x"
	y int "this is y"
}

func main() {
	test_demo1()
	fmt.Println()

	p1 := NewPoint(-1, 2)
	fmt.Println(p1)
	p2 := NewPoint(2, 2)
	fmt.Println(p2)
	fmt.Println()

	p := Point1{1, 2}
	fmt.Println(p)
	pType := reflect.TypeOf(p)
	pField := pType.Field(0)
	fmt.Println(pType)
	fmt.Println(pField)
	fmt.Println("%v\n", pField.Tag) // this is x
}
