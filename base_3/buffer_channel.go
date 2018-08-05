package main

import "fmt"

// 无缓冲信道的发送和接受过程是阻塞的，可以创建一个有缓冲信道，只有在缓冲已满的情况，
// 才会阻塞想缓冲信道发送数据，同样只有在缓冲为空的时候，才会阻塞从缓冲信道接受数据

func main() {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}
