package main

import (
	"fmt"
	"sync"
)

// 互斥锁
type SafeInt struct {
	sync.Mutex // 匿名字段，SafeInt拥有Lock(),Unlock()两个方法
	Num        int
}

func main() {
	count := SafeInt{}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			count.Lock()
			count.Num += i
			fmt.Print(count.Num, " ")
			count.Unlock()
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
