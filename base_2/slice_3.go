package main

import "fmt"

func main() {
	sliceTest()
	fmt.Println()
	twoDimensionArray()
}

func sliceTest() {
	arr := []int {1, 2, 3, 4, 5}
	s := arr[:]
	for e := range s {
		fmt.Printf("%d ", s[e])
	}
	fmt.Println()

	for e := range arr {
		fmt.Printf("%d ", arr[e])
	}
	fmt.Println()

	s1 := make([]int, 3)
	for e := range s1 {
		fmt.Printf("%d ", s1[e])
	}
	fmt.Println()
}

func twoDimensionArray() {
	// 数组
	var a = [][]int {{0, 0}, {1, 2}, {2}, {3, 6}, {4, 8}}
	var i, j int

	// print
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
