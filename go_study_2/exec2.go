package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func exec2() {
	// 在我们的例子中，我们将执行 ls 命令。Go 需要提供我们需要执行的
	// 可执行文件的绝对路径，所以我们将使用exec.LookPath 来得到它
	// （大概是 /bin/ls）
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	fmt.Println(binary)

	// //Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。
	// 我们给 ls 一些基本的参数。注意，第一个参数需要是程序名
	args := []string{"ls", "-a", "-l", "-h"}
	// Exec 同样需要使用环境变量。这里我们仅提供当前的环境变量
	env := os.Environ()

	// 这里是 os.Exec 调用。如果这个调用成功，那么我们的进程将在这里被替换成
	// /bin/ls -a -l -h 进程。如果存在错误，那么我们将会得到一个返回值
	if execErr := syscall.Exec(binary, args, env); execErr != nil {
		panic(execErr)
	}
}
