
// 单元测试的代码文件的名字格式必须是:*_test.go
// 定义函数的时候名字也需要以Test开头

// 开头的calc的名字并不是强制的，但是为了方便测试
// 哪个代码文件，开头就以哪个文件开头
package calc

import (
    "testing"
)

func TestAdd(t *testing.T){
    var sum int
	pipe := make(chan int, 1)
    Add(5, 6, pipe)
	sum = <- pipe
    if sum != 11{
       t.Fatalf("add is not right,sum:%v expected:11",sum)
    }
    t.Logf("add is Ok")
}