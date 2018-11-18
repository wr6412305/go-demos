package main

import (
	"fmt"
)

// error只是一个接口，我们可以自己实现这个接口
/*type error interface {
    Error() string
}*/

type myError struct {
	Desc string
}

func (myErr myError) Error() string {
	return myErr.Desc
}

func doSomething() error {
	return myError{"Logic Error!"}
}

func error1() {
	if err := doSomething(); err != nil {
		fmt.Println(err)
	}
}
