package main

import (
	"fmt"
)

type Ustr string

func (us Ustr) String() string {
	return string(us) + " 自定义格式"
}

func (us Ustr) GoString() string {
	return string(us) + " Go 格式"
}

func (us Ustr) Format(f fmt.State, c rune) {
	switch c {
	case 'm', 'M':
		f.Write([]byte(us + "\n扩展标记:["))
		if f.Flag('-') {
			f.Write([]byte(" -"))
		}
		if f.Flag('+') {
			f.Write([]byte(" +"))
		}
		if f.Flag('#') {
			f.Write([]byte(" #"))
		}
		if f.Flag(' ') {
			f.Write([]byte(" space"))
		}
		if f.Flag('0') {
			f.Write([]byte(" 0"))
		}
		f.Write([]byte(" ]\n"))
		if w, wok := f.Width(); wok {
			f.Write([]byte("宽度值:" + fmt.Sprint(w) + "\n"))
		}
		if p, pok := f.Precision(); pok {
			f.Write([]byte("精度值：" + fmt.Sprint(p)))
		}

	case 'v': // 如果使用 Format 函数，则必须自己处理所有格式，包括 %#v
		if f.Flag('#') {
			f.Write([]byte(us.GoString()))
		} else {
			f.Write([]byte(us.String()))
		}
	default: // 如果使用 Format 函数，则必须自己处理默认输出
		f.Write([]byte(us.String()))
	}
}

func main() {
	us := Ustr("Hello World")
	fmt.Printf("% 0-+#8.5m\n", us)
	fmt.Println(us)
	fmt.Printf("%#v\n", us)
}
