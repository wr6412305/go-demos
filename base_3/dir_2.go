package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetAllFiles(pathname string) (err error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return
	}

	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"/"+fi.Name())
			GetAllFiles(pathname + "/" + fi.Name() + "/")
		} else {
			fmt.Println(fi.Name())
		}
	}

	return
}

func main() {
	var dirPath string
	if os.Args[1] != "" {
		dirPath = os.Args[1]
	}

	GetAllFiles(dirPath)
}
