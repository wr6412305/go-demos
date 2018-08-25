package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readAll(path string) []string {
	var all_file []string
	finfo, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}
	}

	for _, x := range finfo {
		real_path := path + "/" + x.Name()
		all_file = append(all_file, real_path)
		if x.IsDir() {
			all_file = append(all_file, readAll(real_path)...)
		}
	}

	return all_file
}

func main() {
	var dirPath string
	if os.Args[1] != "" {
		dirPath = os.Args[1]
	}

	all_file := readAll(dirPath)
	for _, data := range all_file {
		fmt.Println(data)
	}
}
