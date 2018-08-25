package main

import (
	"fmt"
)

// iota在const关键字出现时将被重置为0(const内部的第一行之前)，
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)

const a = iota // a = 0

const (
	b = iota // b = 0
	c        // c = 1
	d        // d = 2
)

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1<<0 = 00000001
	IgChocolate                         // 1<<1 = 00000010
	IgNuts                              // 1<<2 = 00000100
	IgStrawberries                      // 1<<3 = 00001000
	IgShellfish                         // 1<<4 = 00010000
)

// iota 在下一行增长，而不是立即取得它的引用
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func main() {
	fmt.Printf("%d %d %d %d\n", a, b, c, d)
	fmt.Println()

	fmt.Printf("%d %d %d %d %d\n", IgEggs, IgChocolate,
		IgNuts, IgStrawberries, IgShellfish)
	fmt.Println()

	fmt.Printf("Apple: %d\n", Apple)
	fmt.Printf("Banana: %d\n", Banana)
	fmt.Printf("Cherimoya: %d\n", Cherimoya)
	fmt.Printf("Durian: %d\n", Durian)
	fmt.Printf("Elderberry: %d\n", Elderberry)
	fmt.Printf("Fig: %d\n", Fig)
	fmt.Println()
}
