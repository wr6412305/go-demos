package main

import "fmt"

// 被defer的函数在return之后执行,这个时机点正好可以捕获函数抛出的
// panic,因而defer的另一个重要用途就是执行recover
// recover只有在defer中使用才更有意义,如果在其他地方使用,由于程序
// 已经调用结束而提前返回而无法有效捕捉错误

func main()  {
	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println("recover")
		}
	}()

	panic("error")
}
