package main

import (
		"fmt"
	"reflect"
)

// 接口类型变量转换为反射类型对象

func main(){
	var circle float64 = 6.28
	var icir interface{}

	icir = circle
	fmt.Println("Reflect: circle.Value =", reflect.ValueOf(icir))
	fmt.Println("Reflect: circle.Type =", reflect.TypeOf(icir))

	// 可以看到ValueOf和TypeOf的参数都是空接口，因此，这说明可以直接使用变量传进去
	fmt.Println("Reflect: circle.Value =", reflect.ValueOf(circle))
	fmt.Println("Reflect: circle.Type =", reflect.TypeOf(circle))
}
