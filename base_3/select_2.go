package main

import (
	"time"
	"fmt"
)

// 假设我们有一个关键性应用，需要尽快地把输出返回给用户。这个应用的数据库
// 复制并且存储在世界各地的服务器上。假设函数 server1 和 server2 与这样不
// 同区域的两台服务器进行通信。每台服务器的负载和网络时延决定了它的响应时
// 间。我们向两台服务器发送请求，并使用 select 语句等待相应的信道发出响应
// select 会选择首先响应的服务器，而忽略其它的响应。使用这种方法，我们可
// 以向多个服务器发送请求，并给用户返回最快的响应了

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <- ch:
			fmt.Println("received value:", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
