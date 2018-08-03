package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	fmt.Println(Books{"Go 语言", "liangjisheng", "Go", 12345})
	// 也可以使用键值格式
	fmt.Println(Books{title: "Go 语言", author: "liangjisheng", subject: "Go", book_id: 12345})
	// 忽略的字段为0或空
	fmt.Println(Books{title:"Go 语言", author: "liangjisheng"})
}

