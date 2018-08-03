package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("变量地址: %x\n", &a)

	var ptr *int
	fmt.Printf("ptr的值为: %x\n", ptr)
}
