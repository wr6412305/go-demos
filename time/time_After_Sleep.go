package main

import (
	"fmt"
	"time"
)

// After(d Duration) <- chan Time
// 表示多少时间之后，但是在取出channel内容之前不阻塞，后续程序可以继续执行
//
// Sleep(d Duration)
// 表示休眠多长时间，休眠时处于阻塞状态，后续程序无法执行

func main() {
	fmt.Println("hello")
	// 此处不阻塞1秒，而是直接往下执行
	time.After(time.Second * 1)
	fmt.Println("world")

	fmt.Println("hello")
	// 此处阻塞1秒
	time.Sleep(time.Second * 1)
	fmt.Println("world")
}
