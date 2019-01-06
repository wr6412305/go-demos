package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signal1() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify注册这个给定的通道用于接受特定信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
