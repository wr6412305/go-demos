package main

import "fmt"

// 可以把一个双向信道转换成唯送信道或者唯收信道，但反过来不行

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)
}
