package main

import "fmt"

// 含有defer得语句的函数，会在该函数将要返回之前，调用另一个函数

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

// defer语句不仅限于函数的调用，调用方法也是合法的
type person struct {
	firstname string
	lastname  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstname, p.lastname)
}

func main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
	fmt.Println()

	p := person{
		firstname: "John",
		lastname: "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}
