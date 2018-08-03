package main

import "fmt"

// 未初始化的map的值是nil
func test1(){
	map1 := make(map[string]string, 5)
	map2 := make(map[string]string)
	map3 := map[string]string{}
	map4 := map[string]string{"a":"1", "b":"2", "c":"3"}
	fmt.Println(map1, map2, map3, map4)
	fmt.Println()
}

func test2(){
	ages01 := map[string]int{
		"alice":31,
		"bob":13,
	}
	ages02 := make(map[string]int)
	ages02["chris"] = 20
	ages02["paul"] = 30
	// ages01和ages02两种初始化的方式等价

	//m1和m2创建方式等价，都是创建了一个空的的map，这个时候m1和m2没有任何元素
	m1 := make(map[string]int)
	m2 := map[string]int{}

	for name, age := range ages01{
		fmt.Println(name, "\t", age)
	}
	for name, age := range ages02{
		fmt.Println(name, "\t", age)
	}

	var null_map map[string]int			// 声明但未初始化map，此时是map的零值状态(只有一个nil元素)
	empty_map := map[string]int{}		// 创建了初始化了一个空的的map，这个时候empty_map没有任何元素
	fmt.Println(m1 != nil && m2 != nil)		// true
	fmt.Println(len(null_map) == 0)
	fmt.Println(null_map == nil)	//true,此时是map的零值状态(nil)
	fmt.Println(len(empty_map) == 0)
	fmt.Println(empty_map == nil)	// false,空的的map不等价于nil(map的零值状态)
	empty_map["test"] = 12          //执行正常，空的的map可以赋值设置元素
	// null_map["test"] = 12           //panic: assignment to entry in nil map，无法给未初始化的map赋值设置元素
	fmt.Println()
}

func traverse_map(){
	// range for可用于遍历map 中所有的元素，不过需要注意因为 map本身是无序的，
	// 因此对于程序的每次执行，不能保证使用 range for 遍历 map的顺序总是一致的
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("All items of a map")
	for key, value := range personSalary{
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
}

func main(){
	test1()
	test2()
	traverse_map()
}
