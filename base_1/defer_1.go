package main

// Golang引入关键字defer来延迟执行defer后面的函数.一直等到包含
// defer语句的函数执行完毕时,延迟函数(defer后的函数或语句)才会
// 被执行,而不管包含defer语句的函数是通过return的正常结束,还是
// 由于panic导致的异常结束.你可以在一个函数中执行多条defer语句
// 它们的执行顺序与声明顺序相反
// 有defer关键字之后,即便函数抛出了异常,也会被执行,这样就不会
// 因程序出现了错误而导致资源不会释放了

import "fmt"

func main(){
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
