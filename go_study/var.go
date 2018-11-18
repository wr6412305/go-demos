package main

import (
	"fmt"
	"unsafe"
)

func myvar() {
	var age int // variable declaration
	fmt.Println("my age is ", age)
	age = 29
	fmt.Println("my age is ", age)
	age = 54
	fmt.Println("my age is ", age)

	var age1 int = 10
	fmt.Println("my age1 is ", age1)

	var age2 = 10
	fmt.Println("my age2 is ", age2)

	name, age3 := "ljs", 29
	fmt.Println("my name is", name, "age is", age3)
}

func varSize() {
	a, b := 89, 95
	fmt.Println("a is", a, "b is", b)
	fmt.Printf("type of a is %T, size of a is %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("type of a is %T, size of a is %d", b, unsafe.Sizeof(b))
}

func complexTest() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	var a = 5.9 / 8
	fmt.Printf("a's type %T value %v", a, a)
}
