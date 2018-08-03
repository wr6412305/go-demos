package main

import (
	"fmt"
	)

func main() {
	// 字符串常量没有类型，但有一个相关联的默认类型
	const hello = "hello, world"
	fmt.Printf("type %T value %v\n", hello, hello)
	const typehello string = "HELLO"
	fmt.Printf("type %T value %v\n", typehello, typehello)

	// bool常量
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst
	// defaultBool = customBool		// error
	fmt.Println(defaultBool)
	fmt.Println(customBool)

	// 数字常量
	const a = 5
	fmt.Println(a)
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomples64Var", complex64Var)

}
