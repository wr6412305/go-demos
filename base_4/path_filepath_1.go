package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	s := `https://www.wangshubo.com/a/b/c/d`
	u, _ := url.Parse(s)
	s = u.Path
	fmt.Println(s)
	// FromSlash()将path中的‘/’转换为系统相关的路径分隔符
	s = filepath.FromSlash(s)
	fmt.Println(s)

	// 创建目录
	if err := os.MkdirAll(s[1:], 0777); err != nil {
		fmt.Println(err)
	}

	// ToSlash(path string) // 将path中平台相关的路径分隔符转换为'/'
	s = filepath.ToSlash(s)
	fmt.Println(s)
	fmt.Println()

	path := `a///b///c///d`
	fmt.Println(path)
	path = filepath.FromSlash(path) // 转换为平台相关分隔符
	fmt.Println(path)

	// 返回最后一个分隔符之前的部分(不包含分隔符)
	d1 := filepath.Dir(path)
	fmt.Println(d1)
	// 返回最后一个分隔符之后的部分(不包含分隔符)
	f1 := filepath.Base(path)
	fmt.Println(f1)

	// 获取 path 中最后一个分隔符前后的两部分,之前包含分隔符，之后不包含分隔符
	// Split(path string) (dir, file string)
	// 获取路径字符串中的文件扩展名
	// Ext(path string) string
	d2, f2 := filepath.Split(path)
	fmt.Println(d2)
	fmt.Println(f2)

	ext := filepath.Ext(path)
	fmt.Println(ext)

	// func Join(elem ...string) string
	// 将 elem 中的多个元素合并为一个路径，忽略空元素，清理多余字符
	fmt.Println("On Windows:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))
	fmt.Println()

	// func Abs(path string) (string, error) // 获取 path 的绝对路径
	// IsAbs() 判断路径是否为绝对路径
	s1 := `a/b/c/d`
	fmt.Println(filepath.Abs(s1))
	s2 := `C:\Users\xng\go\src\mycode`
	fmt.Println(filepath.IsAbs(s1)) // false
	fmt.Println(filepath.IsAbs(s2)) // true
}
