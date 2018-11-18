package main

import (
	"fmt"
)

func switch1() {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("incorrect finger number")
	}

	// switch中的表达式式可以省略的，如果省略表达式，则相当于switch true
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

	// Go中执行完一个case后会立刻退出switch语句，fallthrough语句用于表明执行完当前case
	// 语句后按顺序执行下一个case语句
	switch {
	case num < 50:
		fmt.Println("num is greater than 0 and less than 50")
		fallthrough
	case num < 100:
		fmt.Println("num is greater than 51 and less than 100")
		fallthrough
	case num < 200:
		fmt.Println("num is less than 200")
	}
}
