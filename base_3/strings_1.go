package main

import "fmt"
import s "strings" // 给strings取个别名

// 给fmt.Println取一个简单的别名
var p = fmt.Println

func main() {
	// 判断一个字符串是否包含在另一个字符串
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))           // 字符串第一次出现的位置
	p("Join:", s.Join([]string{"a", "b"}, "-")) // 将切片连接
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1)) // 替换所有
	p("Replace:", s.Replace("foo", "o", "0", 1))  // 替换第一个
	p("split:", s.Split("a-b-c-d-e", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
	p()
	p("Len:", len("hello"))
	p("Char:", "hello"[1])
}
