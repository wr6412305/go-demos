package main

import "fmt"

// slice作为参数传递过程中使用的是引用传递

func main(){
	a := []int{1, 2, 3, 4, 5}
	var b = a[0:3]
	var c = [...]int{3, 6, 9, 2, 6, 4}
	d := c[0:2]
	sliceInfo(b)

	fmt.Println("sum of b is", sum(b))
	fmt.Println("sum of d is", sum(d))
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++{
		s += a[i]
	}
	return s
}

func sliceInfo(x []int){
	fmt.Println("len is", len(x), "cap is", cap(x), "slice is", x)
}
