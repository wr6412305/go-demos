package main

import (
	"encoding/json"
	"fmt"
)

// Json Marshal：将数据编码成json字符串

// sex没有导出，所以不能转成json
// 如果变量打上了json标签，如Name旁边的 `json:"name"` ，那么转化成的
// json key就用该标签“name”，否则取变量名作为key，如“Age”，“HIgh”
// 指针变量编码是自动转换为它所指向的值
type Stu struct {
	Name  string `json:"name"`
	Age   int
	High  bool
	sex   string
	Class *Class `json:"class"`
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
