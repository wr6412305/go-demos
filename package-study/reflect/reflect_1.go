package main

import (
	"fmt"
	"reflect"
)

type sr struct {
	string
}

func (s sr) Read() {

}

func (s *sr) Write() {

}

func GetMembers(i interface{}) {
	// 获取i的类型信息
	t := reflect.TypeOf(i)

	for {
		// 进一步获取i的类别信息
		if t.Kind() == reflect.Struct {
			// 只有结构体字段可以获取其字段信息
			fmt.Printf("\n%-8v %v个字段:\n", t, t.NumField())
			// 进一步获取i的字段信息
			for i := 0; i < t.NumField(); i++ {
				fmt.Println(t.Field(i).Name)
			}
		}

		// 任何类型都可以获取其方法信息
		fmt.Printf("\n%-8v %v个方法:\n", t, t.NumMethod())
		// 进一步获取i的方法信息
		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println(t.Method(i).Name)
		}

		if t.Kind() == reflect.Ptr {
			// 如果是指针，获取其所指向的元素
			t = t.Elem()
		} else {
			// 否则上面已经处理过了，直接退出循环
			break
		}
	}
}

func main() {
	GetMembers(&sr{})
}
