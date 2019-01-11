package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	// 计数器增加
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			// 计数器减一
			defer wg.Done()
			fmt.Print(i, " ")
		}(i)
	}

	// 等待直到计数器为0
	wg.Wait()
}
