package main

import (
	"fmt"
	"log"
)

// Golang's log模块主要提供了3类接口。分别是 “Print 、Panic 、Fatal ”，
// 对每一类接口其提供了3中调用方式，分别是 "Xxxx 、 Xxxxln 、Xxxxf"，
// 基本和fmt中的相关函数类似

func test_log1() {
	arr := []int{2, 3}
	log.Print("Print array ", arr, "\n")
	log.Println("Println array", arr)
	log.Printf("Printf array with item [%d, %d]\n", arr[0], arr[1])

	// log.Fatal()先将日志打印到标准输出，接着调用系统的os.Exit(1)接口
	// 由于直接调用os.Exit(1),所以defer函数不会被调用
	defer func() {
		fmt.Println("--first--")
	}()
	log.Fatalln("test for defer Fatal")
}

func test_deferpanic() {
	// 对于log.Panic接口，函数把日志内容刷到标准错误后调用panic函数
	defer func() {
		fmt.Println("--first--")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	log.Panicln("test for defer panic")
	defer func() {
		fmt.Println("--second--")
	}()
}

func main() {
	test_deferpanic()
	fmt.Println()

	test_log1()
}
