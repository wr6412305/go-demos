package main

import "fmt"

// 还可以创建使用指针接收器的方法。值接收器和指针接收器之间的区别在于，
// 在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的

type Employee struct {
	name string
	age  int
}

// 使用值接收器的方法
func (e Employee) changeName(newName string) {
	e.name = newName
}

// 使用指针接收器的方法
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age: 20,
	}
	fmt.Printf("Employee name before change: %s\n", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("Employee name after change: %s\n", e.name)

	fmt.Printf("\nEmployee age before change: %d\n", e.age)
	e.changeAge(21)
	fmt.Printf("Employee age after change: %d\n", e.age)
}
