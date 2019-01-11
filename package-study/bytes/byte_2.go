package main

import (
	"bytes"
	"fmt"
)

func main() {
	b1 := []byte("Hello World!")
	b2 := []byte("Hello 世界！")
	buf := make([]byte, 6)
	// 将b1包装成bytes.Reader对象
	// bytes.Reader 实现了如下接口：
	// io.ReadSeeker
	// io.ReaderAt
	// io.WriterTo
	// io.ByteScanner
	// io.RuneScanner
	rd := bytes.NewReader(b1)
	rd.Read(buf)
	fmt.Printf("%q\n", buf)
	rd.Read(buf)
	fmt.Printf("%q\n", buf)

	// Reset()将底层数据切换为b2，同时复位所有标记(读取位置等信息)
	rd.Reset(b2)
	rd.Read(buf)
	fmt.Printf("%q\n", buf)
	// Len()返回未读取部分的数据长度
	// Size()返回底层数据的总长度，返回值永远不变
	fmt.Printf("Size:%d, Len:%d\n", rd.Size(), rd.Len())
}
