package main

import (
	"fmt"
	"math"
)

// 使用结构体类型和字段提供错误的更多信息
// 创建一个实现error接口的结构体类型，并使用它的字段来提供关于错误的更多信息

// 创建一个表示错误的结构体类型
type areaError struct {
	err string
	radius float64
}

// 一个结构体实现了Error方法就可以从来表示错误
func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle %0.2f", area)
}
