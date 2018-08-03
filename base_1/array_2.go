package main

import "fmt"

// 数组在使用的过程中都是值传递，将一个数组赋值给一个新变量或
// 作为方法参数传递时，是将源数组在内存中完全复制了一份，而不
// 是引用源数组在内存中的地址

func modify_arr(array [5] int){
	array[0] = 10
	fmt.Println("In modify_arr, array values is:", array)
}

func main(){
	arr := [5]int{1, 2, 3, 4, 5}
	modify_arr(arr)
	fmt.Println("In main, array values is:", arr)
}
