package concurrent

import(
	"fmt"
	"time"
)

func test_print(a int){
	fmt.Println(a)
}

func Concurrent(){
	for i:= 0; i < 100; i++ {
		// 在调用函数前面加上go，表示开启了并发
		go test_print(i)
	}
	
	// 表示等待线程结束
	time.Sleep(time.Second)
}