package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	chanN := make(chan bool)
	chanC := make(chan bool, 1)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 11; i += 2 {
			<-chanC
			fmt.Print(i)
			fmt.Print(i + 1)
			chanN <- true
		}
	}()

	go func() {
		charSeq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 0; i < 10; i += 2 {
			<-chanN
			fmt.Print(charSeq[i])
			fmt.Print(charSeq[i+1])
			chanC <- true
		}

		done <- struct{}{}
	}()

	chanC <- true
	<-done
}
