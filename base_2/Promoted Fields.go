package main

import "fmt"

// 如果结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段
// 因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问

type Address struct {
	city, state string
}

type Person struct {
	name string
	age int
	Address		// 匿名结构体类型字段，访问city和state这两个字段就像在Person里直接
				// 声明一样，因此称为提升字段
}

func main(){
	var p Person
	p.name = "Mike"
	p.age = 50
	p.Address = Address{
		city : "Chicago",
		state : "Illinois",
	}

	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("City:", p.city)		// city is promoted field
	fmt.Println("State:", p.state)	// state is promoted field
}
