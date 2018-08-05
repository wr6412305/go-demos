package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Println()
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Printf("s: len = %d, cap = %d\n", len(s), cap(s))
	fmt.Println("address welcome:", &s[0])
	fmt.Println(s)
}

func main() {
	// 把可变参数转换为一个新的切片
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	// 此时nums是长度和容量为0的nil切片
	find(87)

	nums := []int{89, 90, 95}
	// find第二个函数为可变参数，接受int类型的可变参数，但现在nums是一个切片
	// 所以类型错误，
	// find(89, nums)
	// 有一个语法糖，可以在切片后加上...后缀，这样做切片将直接传入函数，不在创建新的切片
	find(89, nums...)

	welcome := []string {"hello", "world"}
	change(welcome...)
	fmt.Printf("welcome: len = %d, cap = %d\n", len(welcome), cap(welcome))
	fmt.Println("address welcome:", &welcome[0])
	fmt.Println(welcome)
}
