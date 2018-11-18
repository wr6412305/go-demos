package main

import (
	"fmt"
)

// i.(T)用来获取接口i的实际类型T的值
// v, ok := i.(T)
// 如果i的具体类型是T，则v将具有i的实际值，ok为true
// 如果i的具体类型不是T,则ok将为false, v将具有T的0值,程序不会触发panic
func assert(i interface{}) {
	// s := i.(int)
	// fmt.Println(s)

	v, ok := i.(int)
	fmt.Println(v, ok)
}

func interface5() {
	var s interface{} = 56
	assert(s)
	s = "ljs"
	assert(s)
}

// 类型断言type switch
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am a int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func interface6() {
	findType("ljs")
	findType(2)
	findType(89.99)
}
