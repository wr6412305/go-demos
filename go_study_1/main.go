package main

import (
	"flag"
	"fmt"
)

func myfmt() {
	o := 0123
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xABC)
	fmt.Printf("%d %[1]x %#[1]x\n", x)
}

func main() {
	var src string
	flag.StringVar(&src, "src", "", "source file")
	var level *int
	level = flag.Int("level", 0, "debug level")
	var memo string
	flag.StringVar(&memo, "memo", "", "the memeory")
	flag.Parse()
	// flag.Usage()
	fmt.Printf("src=%s, level=%d, memo=%s\n", src, *level, memo)

	// myfmt()
	// fieldTag()
	// recover1()
	// for1()
	// for2()
	// for3()
	// for4()
	// reflect1()
	// unsafe1()
	// unsafe2()
	// filepath1("D:\\books")
	// ring1()
	// josephus()
	// list1()
	// list2()
	// sort1()
	// sort2()
	// search()
	// sort3()
	// zip1("F:\\document", "D:\\doc.zip")
	// routine()
	// hash()
	// template1()
	error1()
}
