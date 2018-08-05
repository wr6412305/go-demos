package main

import (
	"sync"
	"fmt"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()	// 引用计数减1
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)		// 引用计数加1
		go increment(&w, &m)	// 创建1000个协程，每个协程对x加1
	}
	w.Wait()	// 等待引用计数为0
	fmt.Println("final value of x", x)
}
