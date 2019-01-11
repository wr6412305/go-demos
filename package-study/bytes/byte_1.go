package main

import (
	"bytes"
	"fmt"
)

func main() {
	bs := [][]byte{
		[]byte("Hello World!"),
		[]byte("Hello 世界！"),
		[]byte("hello golang."),
	}

	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("!！."), r)
	}

	for _, b := range bs {
		// 删除b两边符合条件的字符
		fmt.Printf("%q\n", bytes.TrimFunc(b, f))
	}

	for _, b := range bs {
		fmt.Printf("%q\n", bytes.TrimPrefix(b, []byte("Hello ")))
	}
}
