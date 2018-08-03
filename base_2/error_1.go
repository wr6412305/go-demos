package main

import (
	"errors"
	"math"
	"fmt"
)

// error接口类型的定义
//type error interface {
//	Error() string
//}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	result, err := Sqrt(-1)

	if err != nil {
		// 处理error时会调用Error方法被调用
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}
}
