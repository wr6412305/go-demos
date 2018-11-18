package main

import (
	"fmt"
)

// map的0值为nil.试图给一个nil map添加元素给会导致运行时错误,因此map必须通过make来初始化
func map1() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)

	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
	// 如果键不存在，会返回值类型的0值
	fmt.Println("Salary of joe is", personSalary["joe"])

	newEmp := "joe"
	value, ok := personSalary[newEmp] // 判断一个键是否存在
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	fmt.Println("All items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	fmt.Println("length is", len(personSalary))
}
