package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	// 创建一个 strings 包中的 Reader 对象
	// 它有三个私有字段：s string、i int64、prevRune int
	sr := strings.NewReader("abcdef")
	fmt.Println(sr)

	p := unsafe.Pointer(sr)
	up0 := uintptr(p)
	// 确定要修改的字段（这里不能用 unsafe.Offsetof 获取偏移量，因为是私有字段）
	if sf, ok := reflect.TypeOf(*sr).FieldByName("i"); ok {
		up := up0 + sf.Offset
		p = unsafe.Pointer(up)
		pi := (*int64)(p)
		*pi = 3
	}

	fmt.Println(sr)
	b, err := sr.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
}
