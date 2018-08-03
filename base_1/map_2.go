package main

import "fmt"

func map_add(){
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000     //增加元素
	personSalary["jamie"] = 15000    //增加元素
	personSalary["mike"] = 9000       //增加元素
	fmt.Println("map before change", personSalary)
	personSalary["mike"] = 10000    //修改元素
	fmt.Println("map after change", personSalary)

	// 删除元素需要使用内置函数delete,该函数根据键来删除一个元素
	// 需要强调delete函数没有返回值
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)
}

func map_find(){
	// 根据键值索引某个元素时,也会返回两个值:索引到的值和本次索引是否成功
	// (这里可能会因为索数值越界或者索引键值有误而导致索引失败)
	ages01 := map[string]int{
		"alice":31,
		"bob":13,
	}

	age, ok := ages01["bo"]		//age才是根据键值索引到的值
	if !ok{
		// //索引失败会返回value的零值，这里是int类型，所以是0
		fmt.Printf("索引失败,bo不是map的键值,此时age=%d\n", age)
	}else {
		fmt.Printf("索引成功, age=%d\n", age)
	}
}

// Go语言中map和slice,func一样,不支持==操作符,就是不能直接比较
// 唯一合法的就是和nil作比较,判断该map是不是零值状态
// 如果想自定义一个函数,来比较两个map是否相等,就可以遍历比较它们
// 的键和值是否完全相等,代码如下
func map_equal(x,y map[string] string) bool{
	if(len(x))!= len(y){
		return false
	}
	for k,xv :=range x{
		if yv,ok:=y[k];!ok||yv!=xv{
			return false
		}
	}
	return true
}

func main(){
	map_add()
	map_find()

	m1 := map[string]string{
		"test": "test",
		"asdf": "asdf",
	}
	m2 := map[string]string{
		"test": "test",
		"asdf": "asdf",
	}
	fmt.Println(map_equal(m1, m2))
}
