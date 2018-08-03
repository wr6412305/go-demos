package main

import "fmt"

// 常规函数
func Add(i int, j int) (int) {
	return i + j
}

// 多指返回函数
func Add_Multi_Sub(i, j int) (int, int, int){
	return i + j, i * j, i - j
}

// 变参函数
func sum(nums ... int){
	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)
}


func main(){
	a, b := 2, 3
	arr := []int{1, 2, 3, 4}
	var c = Add(a, b)
	d, e, f := Add_Multi_Sub(a, b)
	fmt.Println(c, d, e, f)
	sum(arr...)		// 注意传参形式
}
