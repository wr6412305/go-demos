package main

import (
	"fmt"
	"unicode/utf8"
	"bytes"
)

func test_string1(){
	// 声明多个字符串并且赋值
	a, b, v := "hello", "world", "widuu"
	fmt.Println(a, b, v)

	// Go语言字符串是UTF-8字符的一个序列
	c := []byte(a)		// 字符串a转换为字节数组，a并没有改变
	c[0] = 'n'			// 修改字节数组c
	d := string(c)		// 将字节数组c转换为字符串
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
}

func test_string2(){
	// 统计字符和字节
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Println("The number of bytes in string str1 is", len(str1))
	fmt.Println("The number of characters in string str1 is", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n",len(str2))
	fmt.Printf("The number of characters in string str2 is %d",utf8.RuneCountInString(str2))
}

func test_string3(){
	// 使用buffer高效拼接字符串
	a, b := "hello", " world"
	var c string = a + b
	fmt.Println(c)

	// 不过用+这种合并方式效率非常低，每合并一次，都是创建一个新的字符串,
	// 就必须遍历复制一次字符串。
	// Java中提供StringBuilder类(最高效,线程不安全)来解决这个问题。
	// Go中也有类似的机制，那就是Buffer(线程不安全)
	var buffer bytes.Buffer
	for i := 0; i < 100; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())
	// 使用bytes.Buffer来组装字符串，不需要复制，只需要将添加的字符串放在
	// 缓存末尾即可。不过需要强调，Golang源码对于Buffer的定义中并没有任何
	// 关于锁的字段,所以Buffer是线程不安全的
}

func main(){
	test_string1()
	fmt.Println()
	test_string2()
}
