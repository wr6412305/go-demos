package main

import (
	"fmt"
	"reflect"
)

type inf interface {
	Method1()
	Method2()
}

type ss struct {
	a func()
}

func (i ss) Method1() {}
func (i ss) Method2() {}

func test(t reflect.Type) {
	if t.NumMethod() > 0 {
		fmt.Printf("\n--- %s ---\n", t)
		fmt.Println(t.NumMethod())
		fmt.Println(t.Method(0).Type)
		fmt.Println(t.Method(0).Func.String())
	}
}

func reflect2() {
	s := reflect.TypeOf(ss{})
	i := reflect.TypeOf(new(inf)).Elem()

	test(s)
	test(i)
}
