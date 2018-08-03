package main

import (
	"fmt"
	"calc"
	"goroute"
	"hello"
	"pipe"
	"concurrent"
	"return_multi_values"
)

func main(){
	// 使用calc包
	pipe1 := make(chan int, 1)
	calc.Add(100, 200, pipe1)
	res := <- pipe1	
	fmt.Println(res)
	fmt.Println()
	
	// 使用goroute包
	goroute.Test_goroute(2)
	fmt.Println()
	
	hello.Hello()
	fmt.Println()
	
	pipe.Test_pipe()
	fmt.Println()
	
	concurrent.Concurrent()
	fmt.Println()
	
	return_multi_values.Return_multi_values()
	fmt.Println()
}