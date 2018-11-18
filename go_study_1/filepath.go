package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func filepath1(path string) {
	filepath.Walk(path, walkfunc)
}

func walkfunc(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}
