package main

import (
	"reflect"
	"fmt"
)

// 反射就是程序能够在运行时检查变量和值，求出他们的类型
// reflect.Type 表示 interface{} 的具体类型，而 reflect.Value 表示它的
// 具体值。reflect.TypeOf() 和 reflect.ValueOf() 两个函数可以分别返回
// reflect.Type 和 reflect.Value

type order struct {
	ordId int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)		// type表示interface的实际类型
	v := reflect.ValueOf(q)
	k := t.Kind()				// kind表示该类型的特定类型，这里是struct
	fmt.Println("Type", t)
	fmt.Println("Value", v)
	fmt.Println("Kind", k)

	// NumField()方法返回结构体中字段的数量,而Field(i int)方法返回字段i的
	// reflect.Value
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field: %d type %T value %v\n",
				i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	o := order{
		ordId: 456,
		customerId: 56,
	}
	createQuery(o)
	fmt.Println()

	// Int()和String()方法可以分别取出reflect.value作为int64和string
	a := 56
	x := reflect.ValueOf(a)
	fmt.Println(x)
	x1 := reflect.ValueOf(a).Int()
	fmt.Printf("type %T value: %v\n", x1, x1)
	b := "hello"
	y := reflect.ValueOf(b)
	fmt.Println(y)
	y1 := reflect.ValueOf(b).String()
	fmt.Printf("type %T value: %v\n", y1, y1)
}
