package main

import (
	"reflect"
	"fmt"
)

// 用反射进行变量修改
//利用反射修改变量时，首先需要使用CanSet函数确认变量是否是可修改的

func main(){
	var circle float64 = 6.28
	value := reflect.ValueOf(circle)
	fmt.Println("Reflect: value =", value)
	fmt.Println("Settability of value:", value.CanSet())

	value2 := reflect.ValueOf(&circle)
	fmt.Println("Settability of value:", value2.CanSet())

	value3 := value2.Elem()
	fmt.Println("Settability of value:", value3.CanSet())

	value3.SetFloat(3.14)
	fmt.Println("Value of value3:", value3)
	fmt.Println("Value of circle:", circle)
}
