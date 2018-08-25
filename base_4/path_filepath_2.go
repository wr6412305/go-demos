package main

import (
	"fmt"
	"path/filepath"
)

// Match(pattern, name string) (matched bool, err error)
// 判断 name 是否和指定的模式 pattern 完全匹配

func main() {
	filename := "start.txt"
	pattern := "*art*"

	matched, err := filepath.Match(pattern, filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(matched)

	pattern = "*fart*"
	matched, err = filepath.Match(pattern, filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(matched)

	filename = "data123.csv"
	pattern = "data[0-9]*"
	matched, err = filepath.Match(pattern, filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(matched)
}
