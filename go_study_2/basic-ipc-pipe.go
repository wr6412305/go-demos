package main

import (
	"fmt"
	"os/exec"
)

func pipe1() {
	cmdGoEnv := exec.Command("go", "env")
	// cmdGrep := exec.Command("grep", "GOROOT")

	stdoutEnv, envError := cmdGoEnv.StdoutPipe()
	if envError != nil {
		fmt.Println("Error happened about standard output pipe", envError)
		return
	}

	if envError := cmdGoEnv.Start(); envError != nil {
		fmt.Println("Error happened in execution", envError)
		return
	}

	a1 := make([]byte, 1024*2)
	n, err := stdoutEnv.Read(a1)
	if err != nil {
		fmt.Println("Error happened in reading from stdout", err)
		return
	}
	fmt.Printf("Standard output of go env command: %s", a1[:n])

	// a2 := make([]byte, 1024)
	// n = stdoutEnv.Read(a1)
	// fmt.Printf("Standard output of go env command: %s\n", a1[:n])
}
