package pipe

import(
	"fmt"
)

// 函数的第一个字母必须大写，然后才能被别的包使用，
// 若是小写，则别的包不可见
func Test_pipe(){
	// 定义一个管道变量，chan表示管道，int表示管道类型，3表示管道容量
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	
	fmt.Println(len(pipe))
	// 打印pipe对象所在的内存地址
	fmt.Println(pipe)
	
	var t1 int
	var t2 int
	var t3 int
	// 从管道中获取数据,遵循先进先出
	t1 = <- pipe
	t2 = <- pipe
	t3 = <- pipe
	
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
}