package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func main() {
	// fmt.Println(GetCurrPath())
	// numToByte()
	// sliceArrayInit()
	// joinByteSlice()
	// myripemd160()
	// testC()
	// filename()
	// err1()
	// string1()
	// pipe1()
	// Factory()
	// signal1()
	// exec1()
	// exec2()
	// ParseUrl()
	// rand1()
	// json1()
	// regex1()
	// ftm1()
	strings1()
}
