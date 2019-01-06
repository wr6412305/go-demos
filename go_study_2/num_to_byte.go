package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func numToByte() {
	var i1 int64 = 511 // [00000000 00000000 ... 00000001 11111111] = [0 0 0 0 0 0 1 255]
	s1 := make([]byte, 0)
	buf := bytes.NewBuffer(s1)

	// 数字转[]byte，网络字节序为大端字节序
	binary.Write(buf, binary.BigEndian, i1)
	fmt.Println(buf.Bytes())

	// 数字转[]byte，小端
	buf.Reset()
	binary.Write(buf, binary.LittleEndian, i1)
	fmt.Println(buf.Bytes())

	s2 := []byte{6: 1, 7: 255}
	buf = bytes.NewBuffer(s2)
	var i2 int64
	binary.Read(buf, binary.BigEndian, &i2)
	fmt.Println(i2)

	s3 := []byte{255, 1, 7: 0}
	buf = bytes.NewBuffer(s3)
	var i3 int64
	binary.Read(buf, binary.LittleEndian, &i3)
	fmt.Println(i3)
}

func sliceArrayInit() {
	var array1 = [5]int{1, 2, 3}
	fmt.Printf("array1--type:%T\n", array1)
	rangeIntPrint(array1[:])

	// 数组不声明长度
	var array2 = [...]int{6, 7, 8}
	fmt.Printf("array2--type:%T\n", array2)
	rangeIntPrint(array2[:])

	// slice
	var array3 = []int{9, 10, 11, 12}
	fmt.Printf("array3--type:%T\n", array3)
	rangeIntPrint(array3)

	// 仅初始化其中的部分元素
	var array4 = [5]string{3: "Chris", 4: "Ron"}
	fmt.Printf("array4--type:%T\n", array4)
	rangeObjPrint(array4[:])

	// 数组的长度将根据初始化的元素确定
	var array5 = [...]string{3: "Tom", 2: "Alice"}
	fmt.Printf("array5--type:%T\n", array5)
	rangeObjPrint(array5[:])

	// slice 长度将根据初始化的元素确定
	var array6 = []string{4: "Smith", 2: "Alice"}
	fmt.Printf("array6--type:%T\n", array6)
	rangeObjPrint(array6)
}

// 打印整型切片
func rangeIntPrint(array []int) {
	for i, v := range array {
		fmt.Printf("index:%d value:%d\n", i, v)
	}
}

func rangeObjPrint(array []string) {
	for i, v := range array {
		fmt.Printf("index:%d value:%s\n", i, v)
	}
}

func joinByteSlice() {
	b1 := []byte("this is a first string")
	b2 := []byte(" this is a second string")
	var buf bytes.Buffer
	buf.Write(b1)
	buf.Write(b2)
	b3 := buf.Bytes()
	fmt.Println(string(b3))

	str1 := "this is a first string"
	str2 := " this is a second string"
	buf.Reset()
	buf.WriteString(str1)
	buf.WriteString(str2)
	str3 := buf.String()
	fmt.Println(str3)
}
