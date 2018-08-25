package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

// go 语言在 window下执行命令 获取本地ip

func main() {
	ww, err := exec.Command("CMD", "/C", " ping 126.com").Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Println(string(ww))

	ww, err = exec.Command("CMD", "/C", " ipconfig").Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	// 匹配ip地址
	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	fmt.Printf("%q\n", reg.FindAllString(string(ww), -1)[0])
}
