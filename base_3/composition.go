package main

import "fmt"

// Go不支持继承，但它支持组合，通过嵌套结构体进行组合

type author struct {
	firstname string
	lastname  string
	bio 	  string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstname, a.lastname)
}

type post struct {
	title		string
	content		string
	// 匿名字段author，字段中的字段可以在当前结构体中直接访问，
	// 被称为提升字段(firstname, lastname, bio)
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("contents of websites")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
	fmt.Println()

	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}

	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}
