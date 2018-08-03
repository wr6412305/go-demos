package main

import "fmt"

// Go 语言不像其它面相对象语言一样可以写个类，然后在类里面写一堆方法，
// 但其实Go语言的方法很巧妙的实现了这种效果：我们只需要在普通函数前面
// 加个接受者（receiver，写在函数名前面的括号里面），这样编译器就知道
// 这个函数（方法）属于哪个struct了。

type A struct {
	name string
}

// 接受者写者函数名前面的括号里面,相当于为结构体A定义了一个成员方法foo
func (a A)foo(){
	fmt.Println(a.name)
}

func main(){
	a := A{"test"}
	a.foo()
}
