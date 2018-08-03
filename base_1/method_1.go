package main

import "fmt"

// type临时定义一个和int具有同样功能的类型，这个类型不能看做int类型的别名
// 他们属于不同的类型，不能直接相互赋值
type myInt int

func Add(a, b int) {			// 函数
	fmt.Println(a + b)
}

func (a myInt) Add (b myInt){	// 方法
	fmt.Println(a + b)
}

func main(){
	a, b := 3, 4
	var aa, bb myInt = 3, 4
	Add(a, b)
	// aa.Add称作选择子（selector），它为接收者aa选择合适的Add方法
	aa.Add(bb)
}
