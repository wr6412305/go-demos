package main

import "fmt"

type Student struct {
	name string
	age int
	weight float32
	score []int
}

func main(){
	// 使用 new 关键字创建一个指针
	pp := new(Student)
	*pp = Student{"student", 23, 65.0, []int{2, 3, 6}}
	fmt.Printf("stu pp have %d subjects\n", len((*pp).score))
	// Go语言自带隐式解引用
	fmt.Printf("stu pp have %d subjects\n", len(pp.score))
}
