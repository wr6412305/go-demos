package main

import (
	"errors"
	"math"
	"fmt"
)

// error类型实际上是抽象了Error()方法的error接口，Golang使用该接口进行
// 标准的错误处理.error对应源代码如下:
// type error interface {
// 		Error() string
// }

// 一般情况下，将error作为多个返回值中的最后一个
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func main(){
	result, err := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}
}
