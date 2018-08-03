package main

import "fmt"

type Student struct {
	name string
	age int
	weight float32
	score []int
}

func main(){
	// //按照字段顺序进行初始化
	stu01 := Student{"stu01", 23, 55.5, []int{95, 96, 98}}
	// 通过field:value 形式初始化，该方式可以自定义初始化字段的顺序
	stu02 := Student{age:23, weight:55.5, score:[]int{97, 98}, name:"stu02"}
	stu01.age = 25
	fmt.Println(stu01.age)
	fmt.Println(stu02)
}
