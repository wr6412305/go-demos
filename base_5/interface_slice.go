package main

import (
	"fmt"
)

// 如果我们定义了一个interface的变量，那么这个变量里面可以存实现这个
// interface的任意类型的对象

type Human struct{
	name string
	age int
	phone string
}

type Student struct {
	Human 		// 匿名字段
	school string
	loan float32
}

type Employee struct{
	Human 
	company string
	money float32
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

// Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employee 重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}


func main() {
	// err := errors.New("emit macho dwarf: elf header corrupted.")

	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	var i Men

	// i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}