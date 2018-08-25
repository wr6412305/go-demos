package main

import (
	"fmt"
)

func main() {
	Int64Slice := []int64{1, 2, 3, 4, 5}
	str1 := "test"
	v1 := map[string]interface{}{
		str1: Int64Slice,
	}

	fmt.Printf("%T\n", v1)
	fmt.Println(v1)
}
