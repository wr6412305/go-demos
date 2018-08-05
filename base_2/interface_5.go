package main

import "fmt"

// 还可以将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与
// 其实现的接口就可以互相比较

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	findType("hello")
	p := Person{
		name: "Mike",
		age: 25,
	}
	findType(p)
}
