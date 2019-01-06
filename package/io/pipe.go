package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

// 管道没有缓冲区

func pipe1() {
	r, w := io.Pipe()

	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {
			n, err = r.Read(buf)
			r.CloseWithError(errors.New("管道被读取端关闭"))
			fmt.Printf("读取：%d, %v, %s\n", n, err, buf[:n])
		}
	}()

	n, err := w.Write([]byte("Hello World!"))
	fmt.Printf("写入：%d, %v\n", n, err)
}

func pipe2() {
	r, w := io.Pipe()
	// 启用一个例程进行读取
	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {
			n, err = r.Read(buf)
			fmt.Printf("读取：%d, %v, %s\n", n, err, buf[:n])
		}
	}()
	// 主例程进行写入
	n, err := w.Write([]byte("Hello World !"))
	fmt.Printf("写入：%d, %v\n", n, err)

	w.CloseWithError(errors.New("管道被写入端关闭"))
	n, err = w.Write([]byte("Hello World !"))
	fmt.Printf("写入：%d, %v\n", n, err)
	time.Sleep(time.Second * 1)
}
