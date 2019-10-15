package main

import (
	"fmt"
	"sync"
)

func cond4() {
	condition := false // 条件不满足
	var mu sync.Mutex
	cond := sync.NewCond(&mu) // 创建一个条件变量

	// 协程创造条件
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
