package main

import "fmt"

// 切片在内部可以认为是由一个结构体类型表示
//type slice struct {
//	Length int
//	Capacity int
//	ZerothElements *byte
//}

func subtactOne(numbers []int) {
	for i := range numbers{
		numbers[i] -= 2
	}
}

func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos)
	fmt.Println()

	// 多维切片
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
	fmt.Println()
}
