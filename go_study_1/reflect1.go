package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
}

func (bird Bird) Fly(xFrom, yFrom int) (xTo, yTo int) {
	xTo = xFrom + 5
	yTo = yFrom + 5
	fmt.Printf("%s fly from {%d, %d} to {%d, %d}\n", bird.Name,
		xFrom, yFrom, xTo, yTo)
	return
}

func (bird Bird) Sing() {
	fmt.Println(bird.Name + " is singing")
}

func reflect1() {
	bird := Bird{"Gold"}
	v := reflect.ValueOf(bird)
	fmt.Println(v.NumMethod()) // 方法个数

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Println(v.Method(i).Type()) // 打印方法类型
	}
	// Output:
	// func(int, int) (int, int)
	// func()

	m := v.MethodByName("Fly")
	p := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
	r := m.Call(p)
	fmt.Println(r[0], r[1])
}
