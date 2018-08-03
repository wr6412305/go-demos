package main

import "fmt"

// 切片本身不拥有任何数，只是对现有数组的引用
// 切片只是底层数组的一种表示，对切片所做的任何修改都会反映在底层数组中
// 当多个切片共用相同的底层数组时，每个切片所做的更改将反映在数组中

func main() {
	numa := [3]int {78, 79, 80}
	nums1 := numa[:]	// create a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[0] = 101
	fmt.Println("array after modification to slice nums2", numa)
	fmt.Println()

	// 切片的长度是切片中的元素数，切片的容量是从创建切片索引开始的底层数组中元素数
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	// fruitslice是从fruitarray的索引1和2创建的.因此fruitlice的长度为2
	// fruiteslice是从fruitarray的索引1创建的.因此fruitslice的容量是从
	// fruitarray 索引为1开始,也就是说从orange开始,该值是6
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))

	// 切片可以重置其容量
	fruitslice = fruitslice[:cap(fruitslice)]	// re-slicing fruitslice till its capacity
	fmt.Println("after re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
	fmt.Println()

	// make创建一个数组，并返回引用该数组的切片
	i := make([]int, 5, 5)
	fmt.Println(i)
}
