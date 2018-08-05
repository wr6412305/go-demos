package main

import "fmt"

// for range循环用于在一个信道关闭之前，从信道接受数据

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)		// 关闭信道
}

func main() {
	ch := make(chan int)
	go producer(ch)
	// 如果信道关闭，循环会自动结束
	for v := range ch {
		fmt.Println("Received", v)
	}
}
