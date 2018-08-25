package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	// 如果err==nil说明文件或者文件夹存在
	if err == nil {
		return true, nil
	}

	// 如果错误使用os.IsNotExist()判断为true，说明文件或者文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}

	// 如果返回的错误为其他类型，则不确定是否存在
	return false, err
}

func main() {
	fun := "flag_1.go"
	if isExist, err := PathExists(fun); isExist && err == nil {
		fmt.Println(fun + " exist")
	}
}
