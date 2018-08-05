package main

import "fmt"

// defer实参取值，在Go语言中，并非在调用延迟函数的时候才确定实参，
// 而是当执行defer语句的时候，就会对延迟函数的实参进行求值

func printA(a int) {
	fmt.Println("value of a in deferred function.", a)
}

func test_defer(){
	// 当一个函数内多次调用defer时，Go会把defer调用放入到一个栈中，按照后进先出原则顺序执行
	name := "Naveen"
	fmt.Printf("orignal string: %s\n", string(name))
	fmt.Println("reversed string: ")
	for _, v := range name {
		defer fmt.Printf("%c", v)
	}
}

func main() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call.", a)
	fmt.Println()

	test_defer()
}
