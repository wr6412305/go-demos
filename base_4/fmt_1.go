package main

import (
	"fmt"
)

type Sample struct {
	a   int
	str string
}

func main() {
	s := new(Sample)
	s.a = 1
	s.str = "hello"
	fmt.Printf("%v\n", s)   // 值的默认格式
	fmt.Printf("%+v\n", s)  // 添加字段名
	fmt.Printf("%#v\n", *s) // 相应值的Go语法表示
	fmt.Printf("%T\n", *s)  // 类型
	fmt.Printf("%%\n")      // 打印一个%号
	fmt.Println()

	a := 123
	fmt.Printf("%1.2d\n", a)
	b := 1.23
	fmt.Printf("%1.1f\n", b)
	c := "asdf"
	fmt.Printf("%*.*s\n", 1, 2, c)

	str1 := fmt.Errorf("%s%d", "error:", 1)
	fmt.Println(str1)
}
