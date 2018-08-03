package main

import "fmt"

// iota是一个特殊常量，可以认为是一个可以被编译器修改的常量，在每一个const
// 关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota,其所
// 代表的数字会自动加1，可以被用作枚举值

func main(){
	const (
		a = iota	// 0
		b			// 1
		c			// 2
		d = "ha"	// 独立值, iota += 1
		e			// "ha"	iota += 1
		f = 100		// iota += 1
		g			// 100 iota += 1
		h = iota	// 7,恢复计数
		i			// 8
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)
}
