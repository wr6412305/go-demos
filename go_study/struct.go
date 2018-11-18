package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func struct1() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	emp2 := Employee{"Thoms", "Paul", 29, 800}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	fmt.Println("First Name:", emp1.firstName)
	fmt.Println("Last Name:", emp1.lastName)
	fmt.Println("Age:", emp1.age)
	fmt.Printf("Salary: $%d\n", emp1.salary)

	emp3 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp3).firstName)
	fmt.Println("Age:", (*emp3).age)
	fmt.Println("First Name:", emp3.firstName)
	fmt.Println("Age:", emp3.age)
}

// 匿名字段虽然没有名字，但匿名字段的默认名字为类型名
type Person struct {
	string
	int
}

func struct2() {
	p := Person{"Naveen", 50}
	fmt.Println(p)
	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)
}

type Address struct {
	city, state string
}

type Person1 struct {
	name string
	age  int
	// 匿名字段也是一个结构体，现在Address中的字段city和state被称为提阶字段
	// 因为他们就好像直接被声明在Person1里的一样
	Address
}

func struct3() {
	var p Person1
	p.name = "ljs"
	p.age = 24
	p.Address = Address{
		city:  "beijing",
		state: "begjing",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}
