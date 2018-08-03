package main

import (
	"fmt"
	"strings"
	"strconv"
)

func test_string1(){
	str01 := `This is a raw string \n`	// 原生字符串
	str02 := "This is a raw string \n"	// 引用字符串
	fmt.Println("原生字符串和引用字符串的区别")
	fmt.Println(str01)
	fmt.Println("")
	fmt.Println(str02)
	fmt.Println("")

	fmt.Println("+连接字符串")
	var str03 string = str01 + str02
	fmt.Println(str03)
	fmt.Println("")

	var str string = "This is an example of a string"
	fmt.Println("HasPrefix 函数的用法")
	fmt.Println("T/F? Does the string", str, "have prefix Th")
	fmt.Println(strings.HasPrefix(str, "Th"))
	fmt.Println("")
	fmt.Println("Contains 函数的用法")
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))

	fmt.Println("Count 函数的用法")
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", ""))
	fmt.Println("")

	fmt.Println("Index 函数的用法")
	// 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.IndexRune("NLT_abc", 'b'))
	// 不存在返回-1
	fmt.Println(strings.IndexRune("NLT_abc", 's'))
	fmt.Println(strings.IndexRune("我是中国人", '中'))
	fmt.Println("")

	fmt.Println("Join函数的用法")
	s := []string {"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, " "))

	fmt.Println()
	fmt.Println("LastIndex函数的用法")
	fmt.Println(strings.LastIndex("go gopher", "go"))	// 3
	fmt.Println("")

	fmt.Println("Replace函数的用法")
	// 最后一个参数表示替换前2个
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	// -1 表示替换所有的
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println("")

	fmt.Println("Split函数的用法")
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.Split("a man a plan a canal panama", "a "))
	fmt.Println(strings.Split(" xyz ", ""))
	fmt.Println("", "Bernardo O'Higgins")
	fmt.Println("")

	fmt.Println("ToLower函数的用法")
	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println("")

	fmt.Println("strconv.Itoa函数用法")
	var an int = 6
	news := strconv.Itoa(an)
	fmt.Println("The new string is", news)
}

func main()  {
	test_string1()
}
