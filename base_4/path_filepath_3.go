package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	// gopath := os.Getenv("GOPATH")
	// fmt.Printf("[%s/bin]\n", gopath)

	gopath := `C:\Users\xng\go\bin`
	list := getShellScript(gopath)
	for i, p := range list {
		fmt.Printf("[%d:%s===%s]\n", i, path.Dir(p), path.Base(p))
	}
}

func getShellScript(rootpath string) []string {
	list := make([]string, 0, 10)
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		list = append(list, path)
		return nil
	})

	if err != nil {
		fmt.Println("walk error [%v]\n", err)
	}
	return list
}
