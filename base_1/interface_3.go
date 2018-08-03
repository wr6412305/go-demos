package main

import "fmt"

// Go语言空interface(interface{})不包含任何的method，因此所有的类型都实现
// 了空interface，空interface在我们需要存储任意类型的数值的时候相当有用

func main(){
	slice := make([]interface{}, 10)
	map1 := make(map[string]string)
	map2 := make(map[string]int)
	map2["TaskID"] = 1
	map1["Command"] = "ping"
	map3 := make(map[string]map[string]string)
	map3["mapvalue"] = map1
	slice[0] = map2
	slice[1] = map1
	slice[2] = map3
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
}
