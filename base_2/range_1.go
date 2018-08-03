package main

import "fmt"

// 在for循环中迭代数组(array)、切片(slice)、通道(channel)或者集合(map)
// 在数组和切片中返回元素索引和对应的值,在map中返回key-value对

func main() {
	nums := [] int {2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range用在map上
	kvs := map[string] string {"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range可以枚举Unicode字符串,第一个参数为字符的索引,第二个是Unicode字符本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
