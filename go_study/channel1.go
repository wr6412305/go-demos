package main

import (
	"fmt"
)

// v, ok := <- ch
// ok为true表示成功接收数据，false表示信道已经被关闭，得到信道类型的0值
// 发送者可以关闭信道通知接收者将不会在发送数据

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func channel2() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received", v, ok)
	}
}

func channel3() {
	ch := make(chan int)
	go producer(ch)
	// 不断从信道中接收数据，知道信道关闭，循环自动退出
	for v := range ch {
		fmt.Println("Received", v)
	}
}
