package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func filename() {
	fullFilename := "C:\\doc\\test.txt"
	fmt.Println("fullFilename =", fullFilename)

	var filenameWithSuffix string
	filenameWithSuffix = filepath.Base(fullFilename) //获取文件名带后缀
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)

	var fileSuffix string
	fileSuffix = filepath.Ext(filenameWithSuffix) //获取文件后缀
	fmt.Println("fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
	fmt.Println("filenameOnly =", filenameOnly)
}
