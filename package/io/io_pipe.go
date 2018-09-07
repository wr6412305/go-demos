package main

import (
	"errors"
	"fmt"
	"io"
)

// 管道(读取端关闭)

func main() {
	r, w := io.Pipe()

	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {
			n, err = r.Read(buf)
			r.CloseWithError(errors.New("管道被读取端关闭"))
			fmt.Printf("读取: %d, %v, %s\n", n, err, string(buf[:n]))
		}
	}()

	// 主协程写入
	n, err := w.Write([]byte("Hello World!"))
	fmt.Printf("写入:%d, %v\n", n, err)
}
