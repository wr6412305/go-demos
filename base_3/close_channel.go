package main

import "fmt"

// 接收方可以多用一个变量来检查信道是否已经关闭
// v, ok := <- ch

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)		// 关闭信道
}

func main() {
	ch := make(chan int)
	go producer(ch)
	// 无限循环
	for {
		v, ok := <- ch
		// ok为false，说明信道关闭，退出循环
		if ok == false {
			break
		}
		fmt.Println("Recvived", v, ok)
	}
}
