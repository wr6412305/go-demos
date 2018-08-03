package main

import "fmt"

func main(){
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if 0 == i % j {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}
