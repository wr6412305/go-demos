package main

import (
	"reflect"
	"fmt"
)

// 反射类型对象转换为接口类型变量

func main(){
	var circle float64 = 6.28
	var icir interface{}
	icir = circle

	valueref := reflect.ValueOf(icir)
	fmt.Println(valueref)
	fmt.Println(valueref.Interface())

	y := valueref.Interface().(float64)
	fmt.Println(y)
}
