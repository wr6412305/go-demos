package main

import (
	"log"
	"os"
	"text/template"
)

func template1() {
	// 创建一个名为report的模板，并解析一个字符串
	t, err := template.New("report").Parse("I am {{.}} years old.")
	if err != nil {
		log.Fatal(err)
	}

	// 执行模板，5为参数，将结果写入标准输出
	t.Execute(os.Stdout, 5)
}
