package main

import (
	"fmt"

	"github.com/prashantv/gostub"
)

// Cleanup ...
var Cleanup = cleanup

// 没有返回值的函数称为过程 通常将资源清理类函数定义为过程
func cleanup(val string) {
	fmt.Println(val)
}

func stubProcess() {
	stubs := gostub.StubFunc(&Cleanup)
	Cleanup("hello go")
	defer stubs.Reset()
}
