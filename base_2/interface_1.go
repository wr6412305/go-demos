package main

import "fmt"

type myinterface interface {
	name() string
	age() int
}

type Woman struct {

}

func (woman Woman) name() string {
	return "Woman"
}

func (woman Woman) age() int {
	return 23
}

type Men struct {

}

func (men Men) name () string {
	return "Men"
}

func (men Men) age() int {
	return 27
}

func main() {
	var myif myinterface
	myif = new(Woman)

	fmt.Println(myif.name())
	fmt.Println(myif.age())

	myif1 := new(Men)
	fmt.Println(myif1.name())
	fmt.Println(myif1.age())
}
