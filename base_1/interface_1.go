package main

import "fmt"

// 接口定义了一组方法集合，但是这些方法不包含具体的实现代码
// 还有需要强调一点：接口定义中不能包含变量
// 在定义了一个接口之后,一般使用一个自定义结构体(struct)去实现接口中的方法
// 对于某个接口的同一个方法,不同的结构体(struct)可以有不同的实现

type Phone interface {
	call()
}

type NokiaPhone struct {}
type IPhone struct {}

func (nokiaPhone NokiaPhone) call () {
	fmt.Println("I am Nokia, I can call you!")
}

func (iPhone IPhone) call () {
	fmt.Println("I am iPhone, I can call you!")
}

func main(){
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
