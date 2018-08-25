package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// ReadDir 读取指定目录中的所有目录和文件（不包括子目录）
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
