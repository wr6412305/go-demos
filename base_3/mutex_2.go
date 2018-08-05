package main

import (
	"sync"
	"fmt"
)

// 使用信道处理竞态条件

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	// 缓冲信道容量为1,其他协程试图写入该信道时，都会发送阻塞，只允许有一个协程将数据写入信道
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}

	w.Wait()
	fmt.Println("final value of x", x)
}
