package main

import "fmt"

// 不像大多数面向对象编程语言,在Go语言中接口可以有值
// Go语言中接口不仅有值,还能进行接口赋值,接口赋值分为以下两种情况:
// 1.将对象实例赋值给接口,将对象实例赋值给接口要求对象实现了接口的所有方法
// 2.将一个接口赋值给另一个接口, 接口之间的赋值要求接口A
// 中定义的所有方法,都在接口B中有定义,那么B接口的实例可
// 以赋值给A的对象.反之不一定成立,除非A和B定义的方法完全一样
// (顺序不要求),这时A和B等价,可以相互赋值

type Shaper interface {
	Area() float32
}

type Square struct{
	side float32
}

type Cube struct {
	side float32
}

func (sq *Square) Area() float32{
	return sq.side * sq.side
}

func (cu *Cube) Area() float32{
	return cu.side * cu.side * cu.side
}

func main(){
	sq1 := new(Square)
	sq1.side = 5
	cu1 := new(Cube)
	cu1.side = 5

	// 赋值方法1
	var areaIntf Shaper
	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	areaIntf = cu1
	fmt.Printf("The cube has area: %f\n", areaIntf.Area())

	// 更短的赋值方法2
	areaIntf = Shaper(sq1)
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	areaIntf = Shaper(cu1)
	fmt.Printf("The cube has area: %f\n", areaIntf.Area())
}
