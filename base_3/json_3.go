package main

import (
	"encoding/json"
	"fmt"
)

// 有时为了通用性，或使代码简洁，我们希望有一种类型可以接受各种类型的数据，
// 并进行json编码。这就用到了interface{}类型

type Stu struct {
	// 如果变量打上了json标签,则转换为json格式时，使用标签名name
	Name  interface{} `json:"name"`
	Age   interface{}
	High  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	stu := Stu{
		Name: "张三",
		Age:  18,
		High: true,
		sex:  "男",
	}

	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla

	// Marshal失败是err!=nil
	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
		return
	}

	// jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))
}
