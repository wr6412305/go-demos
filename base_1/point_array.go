package main

import "fmt"

const MAX int = 3

func main() {
	a := []int { 10, 100, 1000 }
	var i int
	var ptr [MAX] *int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %x\n", i, ptr[i])
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
