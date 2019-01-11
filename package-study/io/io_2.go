package main

import (
	"bytes"
	"fmt"
	"os"
)

// bytes.NewBuffer 实现了很多基本的接口

func main() {
	buf := bytes.NewBuffer([]byte("Hello World!"))
	b := make([]byte, buf.Len())
	fmt.Printf("%s\n", buf.String())

	n, err := buf.Read(b)
	fmt.Printf("%s %v\n", b[:n], err)
	// Hello World!   <nil>

	buf.WriteString("ABCDEFG\n")
	buf.WriteTo(os.Stdout)
	// ABCDEFG

	fmt.Printf("%s\n", buf.String())
	n, err = buf.Write(b)
	fmt.Printf("%d %s %v\n", n, buf.String(), err)
	// 12   Hello World!   <nil>

	c, err := buf.ReadByte()
	fmt.Printf("%c %s %v\n", c, buf.String(), err)
	// H ello World! <nil>

	c, err = buf.ReadByte()
	fmt.Printf("%c %s %v\n", c, buf.String(), err)
	// e llo World! <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s %v\n", buf.String(), err)
	// ello World! <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s %v\n", buf.String(), err)
	// ello World! bytes.Buffer: UnreadByte: previous operation was not a successful read
}
