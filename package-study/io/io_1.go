package main

import (
	"fmt"
	"io"
)

// 自定义对象实现io包里面的Reader接口

type Ustr struct {
	s string // 数据流
	i int    // 读写位置
}

// 根据字符串创建Ustr对象
func NewUstr(s string) *Ustr {
	return &Ustr{s, 0}
}

// 获取未读取部分的数据长度
func (s *Ustr) Len() int {
	return len(s.s) - s.i
}

// type Reader interface {
// Read(p []byte) (n int, err error)
// }

// 实现Ustr类型的Read方法
func (s *Ustr) Read(p []byte) (n int, err error) {
	for ; s.i < len(s.s) && n < len(p); s.i++ {
		c := s.s[s.i]
		// 将小写字母转换为大写字母，然后写入p中
		if 'a' <= c && c <= 'z' {
			p[n] = c + 'A' - 'a'
		} else {
			p[n] = c
		}

		n++
	}

	if 0 == n {
		return n, io.EOF
	}
	return n, nil
}

// 然后就可以用 ReadFull 方法读取 Ustr 对象的数据了

func main() {
	s := NewUstr("Hello World!")
	buf := make([]byte, s.Len())

	n, err := io.ReadFull(s, buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", string(buf))
	fmt.Println(n, err)
}
