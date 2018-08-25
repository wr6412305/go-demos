package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func PathExists(path string) (bool, error) {
	// 如果返回的错误为nil,说明文件或文件夹存在
	// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	// 如果返回的错误为其它类型,则不确定是否在存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	filename := "test1.txt"
	IsExist, err := PathExists(filename)
	if !IsExist {
		fmt.Println("file not exist")
		return
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	fmt.Print(string(data))
}
