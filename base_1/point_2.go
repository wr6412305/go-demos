package main

import (
	"strings"
	"fmt"
)

type Person struct{
	firstname string
	lastname string
}

func upPerson(p *Person){
	p.firstname = strings.ToUpper(p.firstname)
	p.lastname = strings.ToLower(p.lastname)
}

func main(){
	// 1-struct as a value type:
	var pers1 Person
	pers1.firstname = "Chris"
	pers1.lastname = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstname, pers1.lastname)

	// 2-struct as a pointer:
	pers2 := new(Person)
	pers2.firstname = "Chris"
	pers2.lastname = "Woodward"
	(*pers2).lastname = "Woodward"		// 这是合法的
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstname, pers2.lastname)

	// 3-struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstname, pers3.lastname)
	fmt.Println()
}
