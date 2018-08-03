package main

import "fmt"

func test_switch1(){
	i := 2
	// 每个case块不需要特别使用break语句来表示结束，执行完某个case语句后
	// 会退出整个switch代码块，默认相当于每个case最后带有break
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		// 如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，
		// 可以使用fallthrough关键字来达到目的
		// fallthrough会强制执行后面的case代码
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Default")
	}
}

func test_switch2(){
	Num := 6
	var grade string = "B"
	var marks int = 90
	switch  {
	case 0 <= Num && Num <= 3:
		fmt.Println("0-3")
	case 4 <= Num && Num <= 6:
		fmt.Println("4-6")
	case 7 <= Num && Num <= 9:
		fmt.Println("7-9")
	}

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50, 60, 70: grade = "C"
	default: grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("Excellent!")
	case grade == "B", grade == "C":
		fmt.Println("Well done")
	case grade == "D":
		fmt.Println("You passed")
	case grade == "F":
		fmt.Println("Better try again")
	default:
		fmt.Println("Invalid grade")
	}

	fmt.Println("Your grade is:", grade)
}

func main(){
	test_switch1()
	fmt.Println()
	test_switch2()
}
