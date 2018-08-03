package main

import "fmt"

// 当新的元素被添加到切片时，会创建一个新的数组。
// 现有数组的元素被复制到这个新数组中，并返回这个新数组的新切片引用

func main() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
	fmt.Println()

	// 一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片
	var names []string	// zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
		fmt.Println()
	}

	// 也可以使用 ... 运算符将一个切片添加到另一个切片
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}
