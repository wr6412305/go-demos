package main

import (
	"fmt"
)

func for1() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	var recs []func()
	for i, v := range a {
		// for循环引入了一个新的词块，在这个词块中声明了i, v.所有在循环体中创建的匿名
		// 函数都捕捉了这两个变量本身(变量的地址)，而不是捕捉了这两个变量的值。因此
		// 在试图恢复数据的值时，每个匿名函数都是以第一个for循环结束时i,v(2, 3)
		// 的值带入。相当于执行了3次a[2]=3.
		// 为了解决这个问题，通常的做法是在循环体内部声明一个同名变量的拷贝
		// 将这个新声明的变量带入匿名函数:见for2()
		a[i] = 0
		recs = append(recs, func() {
			a[i] = v
		})
	}

	for _, rec := range recs {
		rec()
	}

	fmt.Println(a)
}

func for2() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	var recs []func()
	for i, v := range a {
		i, v := i, v // 声明同名变量，拷贝i,v的值
		a[i] = 0
		recs = append(recs, func() {
			a[i] = v
		})
	}

	for _, rec := range recs {
		rec()
	}
	fmt.Println(a)
}

func for3() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan struct{})

	// 在闭包里捕获了循环变量。在这个例子里,for循环的每一次迭代都在开启了
	// 一个协程后马上进入下一次迭代，开启另一个协程，这时上一个协程可能还
	// 没来得及运行，但是捕获的循环变量的值却已经改变，导致了多个协程处理
	// 同一个值的情况。解决方法是在闭包里避免捕获循环变量，因为Go函数参数
	// 都是值传递，因此可以给闭包增加一个参数，并将循环变量作为实参传递
	// 见for4()
	for _, v := range s {
		go func() {
			fmt.Printf("%d ", v)
			ch <- struct{}{}
		}()
	}

	// 等待所有协程结束
	for range s {
		<-ch
	}
	fmt.Println()
}

func for4() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan struct{})

	for _, v := range s {
		go func(v int) {
			fmt.Printf("%d ", v) // v是闭包的形参，值与循环变量一致
			ch <- struct{}{}
		}(v) // 循环变量作为实参传入
	}

	for range s {
		<-ch
	}
	fmt.Println()
}
