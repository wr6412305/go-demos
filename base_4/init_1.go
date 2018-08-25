package main

import (
	"fmt"
	"runtime"
)

var (
	prompt = "Enter a digit, e.g.3 " + "or %s to quit."
)

func init() {
	if runtime.GOOS == "windows" {
		fmt.Println("windows")
	} else { // Unix-like
		fmt.Println("Unix-like")
	}
}

func init() {
	fmt.Println("init")
}

func main() {
	fun := "main"

	fmt.Println(fun)
}
