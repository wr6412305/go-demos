package main

import "fmt"

// 无缓冲的信道在取消息和存消息的时候都会挂起当前的goroutine，
// 除非另一端已经准备好。如果不用信道来阻塞主线的话，主线程就会过早跑完，
// loop线程都没有机会执行
// 需要强调的是,无缓冲的信道永远不会存储数据,只负责数据的流通.体现在:
// 1.从无缓冲信道取数据，必须要有数据流进来才可以，否则当前线阻塞
// 2.数据流入无缓冲信道, 如果没有其他goroutine来拿走这个数据，那么当前线阻塞

var complete chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	complete <- 10		// 执行完毕，发个消息
}

func main() {
	go loop()

	nRes := <- complete		// 直到线程跑完,取到消息. main在此阻塞住
	fmt.Println("nRes =", nRes)
}
