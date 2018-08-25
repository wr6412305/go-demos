package main

import (
	"fmt"
	"time"
)

func main() {
	// 返回当前时间截
	nowTime := time.Now().Unix()
	fmt.Println(nowTime)
	// str格式化时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 时间截转str格式化时间
	str_time := time.Unix(nowTime, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str_time)

	// str格式化时间转换为时间截
	the_time := time.Date(2018, 8, 22, 20, 06, 20, 0, time.Local)
	unix_time := the_time.Unix()
	fmt.Println(unix_time)

	t1, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	if err == nil {
		unix_time := t1.Unix()
		fmt.Println(unix_time)
	}
	fmt.Println()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("After 5 second")
	}

	c := time.Tick(2 * time.Second)
	for now := range c {
		fmt.Println(now)
	}
}
