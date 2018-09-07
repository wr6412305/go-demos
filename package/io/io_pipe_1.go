package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

// 管道(写入端关闭)

func main() {
	r, w := io.Pipe()

	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {
			n, err = r.Read(buf)
			fmt.Printf("读取: %d, %v, %s\n", n, err, string(buf[:n]))
		}
	}()

	// 主协程写入
	n, err := w.Write([]byte("Hello World!"))
	fmt.Printf("写入: %d, %v\n", n, err)

	w.CloseWithError(errors.New("管道被写入端关闭"))
	n, err = w.Write([]byte("Hello World!"))
	fmt.Printf("写入: %d, %v\n", n, err)
	time.Sleep(time.Second * 1)
}
