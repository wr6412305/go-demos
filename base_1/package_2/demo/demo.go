package main

import (
	"fmt"
	"config"
)

func main(){
	fmt.Println(config.ReadConfig())
	fmt.Println("port=", config.Port)
}
