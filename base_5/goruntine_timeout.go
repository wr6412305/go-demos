package main

import "time"

func main(){
	c := make(chan int)
	o := make(chan bool)
	go func(){
		for {
			select {
			case v := <- c:
				println(v)
			// 5s之内下面这个分支阻塞，5s后不阻塞，执行这个分支，函数退出
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()

	<- o
}
