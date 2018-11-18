package main

import (
	"fmt"
	"sort"
)

func sort1() {
	is := []int{3, 5, 2}
	fmt.Println(sort.IntsAreSorted(is)) // false
	sort.Ints(is)
	fmt.Println(is)
	fmt.Println(sort.IntsAreSorted(is)) // true
	fmt.Println(sort.SearchInts(is, 5)) // 2
	fmt.Println(sort.SearchInts(is, 8)) // 3
}

// 自定义类型排序通过两种方式，一种是传递函数值，另一种是通过实现接口
// 1.传递函数值进行排序
type Person struct {
	Name string
	Age  int
}

func sort2() {
	persons := []Person{
		Person{"Gopher", 11},
		Person{"Monkey", 12},
		Person{"Cynhard", 5},
	}

	// 按照名字排序
	less := func(i, j int) bool {
		return persons[i].Name < persons[j].Name
	}

	fmt.Println(sort.SliceIsSorted(persons, less)) // false
	sort.Slice(persons, less)
	fmt.Println(persons)
	fmt.Println(sort.SliceIsSorted(persons, less)) // true
}

// 二分法查找，这个函数很有意思，它查找并返回使函数f返回true和返回false的临界点
// sort.Search()
func search() {
	ints := []int{22, 34, 21, 32, 54, 64, 49, 43}
	sort.Ints(ints)
	fmt.Println(ints)

	// ints[0]~ints[1]都小于32，ints[2]~ints[7]都大于等于32
	// 因此临界点索引为2，found==2
	found := sort.Search(len(ints), func(i int) bool {
		return ints[i] >= 32
	})
	fmt.Println(found) // 2

	if found < len(ints) && ints[found] == 32 {
		fmt.Printf("32 found at index %d\n", found)
	} else {
		fmt.Println("32 not found")
	}

	// 由于找不到一个临界点，使序列左边为32，右边不为32
	// 所以返回len(ints)，found==8
	found = sort.Search(len(ints), func(i int) bool {
		return ints[i] == 32
	})
	fmt.Println(found)
}

// 2.实现接口进行自定义类型排序
// sort.Interface，所有满足这个接口的自定义类型都可以进行排序
// type Interface interface {
// Len() int
// Less(i, j int) bool
// Swap(i, j int)
// }
type PerColl []Person

var persons PerColl

func (PerColl) Len() int {
	return len(persons)
}

func (PerColl) Less(i, j int) bool {
	return persons[i].Name < persons[j].Name
}

func (PerColl) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

func sort3() {
	persons = []Person{
		Person{"Cynhard", 5},
		Person{"Gopher", 11},
		Person{"Monkey", 12},
	}

	sort.Sort(persons)
	fmt.Println(persons)
	fmt.Println(sort.IsSorted(persons))
	sort.Sort(sort.Reverse(persons))
	fmt.Println(persons)
}
