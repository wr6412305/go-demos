package main

import "fmt"

// 在Go中,应用程序并发处理的部分被称作goroutines(go协程)

func loop(){
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	// 把一个loop放在一个goroutine里跑，我们可以使用关键字go来定义并启动一个goroutine
	go loop()
	// 有可能上面的goroutine还没来得及跑loop的时候，主函数已经退出了
	// 如何让goroutine告诉主线程我执行完毕了？使用一个信道来告诉主线程即可
	// 代码见Goroutint_2.go

	loop()
}
