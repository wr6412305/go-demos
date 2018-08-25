package main

import (
	"flag"
	"fmt"
)

func main() {
	married := flag.Bool("married", false, "Are you married?")
	age := flag.Int("age", 22, "How old are you?")
	name := flag.String("name", "", "What your name?")

	var address string
	// flag.StringVar这样的函数第一个参数换成了变量地址，后面的参数和flag.String是一样的
	flag.StringVar(&address, "address", "shanxi", "Where is your address?")

	flag.Parse() // 解析输入的参数

	fmt.Println("married:", *married)
	fmt.Println("age:", *age)
	fmt.Println("name:", *name)
	fmt.Println("address:", address)
}
