package main

import (
	"fmt"
	"reflect"
)

// 结构体字段标签
type Point struct {
	X int `desc:"x position" author:"ljs"`
	Y int `desc:"y position"`
}

func fieldTag() {
	pt := Point{1, 2}
	tp := reflect.TypeOf(pt)
	fmt.Println(tp)
	// fmt.Println(reflect.TypeOf(tp))

	if tp.Kind() == reflect.Struct {
		// 通过字段名获得这个字段的所有信息，如果这个字段存在的话
		if fieldX, found := tp.FieldByName("X"); found {
			fmt.Printf("%s created by %s\n", fieldX.Tag.Get("desc"), fieldX.Tag.Get("author"))
		}
	}

	if fieldY, found := tp.FieldByName("Y"); found {
		if desc, found := fieldY.Tag.Lookup("desc"); found {
			fmt.Printf(`desc of Y is "%s"`+"\n", desc)
		}
	}
}
