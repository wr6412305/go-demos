package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true		// 向信道写入数据
}

func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go routine")
	go hello(done)
	<- done		// 从信道获取数据，如果没有数据，主协程将被阻塞
	fmt.Println("Main received data")
}
