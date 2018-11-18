package main

import (
	"fmt"
	"sync"
)

// 将共享变量的访问限制在一个协程中，避免了数据竞争，不过必须要建立一个监听线程
// 来专门处理共享变量的修改

var tickCount = 200 // 总票数
// 用来控制tickCount的同步，用10个缓冲信道个数模拟10个售票退票窗口
var ch = make(chan int, 10)
var n sync.WaitGroup           // 用来等待购票和退票动作完成
var done = make(chan struct{}) // 等待监听协程退出

// 购票
func buy() {
	ch <- -1
}

// 退票
func refund() {
	ch <- 1
}

func routine() {
	// 监听协程
	go func() {
		for amount := range ch {
			tickCount += amount
			fmt.Println("tick count:", tickCount)
			n.Done() // 每次调用Done()，n的计数减1
		}
		done <- struct{}{}
	}()

	n.Add(2)    // 因为要执行两个动作，所以使n的计数加2
	go buy()    // 购票
	go refund() // 退票

	n.Wait() // 等待n的计数为0

	close(ch) // 关闭信道
	<-done    // 等待监听协程结束

	fmt.Println("tick count:", tickCount)
}
