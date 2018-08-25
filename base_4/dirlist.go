package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// 仅仅获取一个目录下的文件和文件夹的列表
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
	}
	fmt.Println()

	// 另外一个仅仅获取一个目录下的文件和文件夹的列表
	filelist, _ := filepath.Glob("*")
	// contains a list of all files in the current directory
	fmt.Println(filelist)
}
