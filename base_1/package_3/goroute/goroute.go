package goroute

import "fmt"

// 函数的第一个字母必须大写，然后才能被别的包使用，
// 若是小写，则别的包不可见
func Test_goroute(a int){
	fmt.Println(a)
}