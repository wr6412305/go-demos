package main

// 可以把接口看做内部的一个元组(type, value),type 是接口底层的具体类型(Concrete Type),
// 而value是具体类型的值

import (
	"fmt"
)

type Test interface {
	Tester()
}

type myFloat64 float64

func (m myFloat64) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Test
	f := myFloat64(89.7)
	t = f
	describe(t)
	t.Tester()
	f.Tester()
}
