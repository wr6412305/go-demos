package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	io.WriteString(os.Stdout, "Hello World!\n")

	r := strings.NewReader("Hello World!")
	b := make([]byte, 15)

	n, err := io.ReadAtLeast(r, b, 20)
	fmt.Printf("%q %d %v\n", string(b[:n]), n, err)

	r.Seek(0, 0)
	b = make([]byte, 15)

	n, err = io.ReadFull(r, b)
	fmt.Printf("%q %d %v\n", string(b[:n]), n, err)
}
