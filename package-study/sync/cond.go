package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func cond1() {
	condition := false // 条件不满足
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	go func() {
		mu.Lock()
		condition = true // 更改条件
		cond.Signal()    // 发送通知：条件已经满足
		mu.Unlock()
	}()

	mu.Lock()
	// 检查条件是否满足，避免虚假通知，同时避免 Signal 提前于 Wait 执行
	for !condition {
		// 等待条件满足的通知，如果收到虚假通知，则循环继续等待
		cond.Wait() // 等待时 mu 处于解锁状态，唤醒时重新锁定
	}

	fmt.Println("条件满足，开始后续动作...")
	mu.Unlock()
}

func cond2() {
	runtime.GOMAXPROCS(4)
	testCond()
}

func testCond() {
	c := sync.NewCond(&sync.Mutex{})
	condition := false

	go func() {
		time.Sleep(time.Second * 1)
		c.L.Lock()
		fmt.Println("[1] 变更condition状态,并发出变更通知.")
		condition = true
		c.Signal() // c.Broadcast()
		fmt.Println("[1] 继续后续处理.")
		c.L.Unlock()
	}()

	c.L.Lock()
	fmt.Println("[2] condition..........1")
	for !condition {
		fmt.Println("[2] condition..........2")
		// 等待Cond消息通知
		c.Wait()
		fmt.Println("[2] condition..........3")
	}
	fmt.Println("[2] condition..........4")
	c.L.Unlock()

	fmt.Println("main end...")
}
