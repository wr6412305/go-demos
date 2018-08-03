package main

import "fmt"

// k = 3 << 2，l = 3 << 3
const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main(){
	fmt.Println("i =", i)
	fmt.Println("j =", j)
	fmt.Println("k =", k)
	fmt.Println("l =", l)
}
