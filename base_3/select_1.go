package main

import (
	"time"
	"fmt"
)

// select用于在多个接受/发送信道操作中进行选择。select语句会一直阻塞，知道接受/发送操作准备就绪
// 如果有多个信道操作准备完毕，select会随机的选取其中之一执行

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	// 程序运行到这里，会阻塞，除非其中有case准备就绪
	select {
	case s1 := <- output1:
		fmt.Println(s1)
	case s2 := <- output2:
		fmt.Println(s2)		// output2信道首先有数据写入，所以这里会走case2分支，然后程序结束
	}
}
