package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	filename := "test.csv"
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	w := csv.NewWriter(file)
	w.Write([]string{"123", "234234", "345345", "234234"})
	w.Flush()
	file.Close()

	rfile, _ := os.Open(filename)
	r := csv.NewReader(rfile)
	strs, _ := r.Read()
	fmt.Println(strs)
	for _, str := range strs {
		fmt.Println(str, "\t")
	}
}
