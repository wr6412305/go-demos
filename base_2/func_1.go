package main

import "fmt"

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return 	// 不需要明确指定返回值，默认返回area, perimeter的值
}

// 命名返回值
func rectProps1(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return 	// 不需要明确指定返回值，默认返回area, perimeter的值
}

func main() {
	area, peri := rectProps(10.8, 5.6)
	fmt.Println(area, peri)
	area, peri = rectProps1(10.8, 5.6)
	fmt.Println(area, peri)
	// 空白符_, 返回值周长被丢弃
	area, _ = rectProps(10.8, 5.6)
	fmt.Println(area)
}
