package main

import "fmt"

func test_array1(){
	myArray := [3][4] int {{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	// 打印一维数组长度
	fmt.Println(len(myArray))
	// 打印二维数组长度
	fmt.Println(len(myArray[1]))
	// 打印整个二维数组
	fmt.Println(myArray)
	fmt.Println()
}

func test_array2(){
	var arr1[5] int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Println("Array at index", i, "is", arr1[i])
	}
	fmt.Println()
}

func test_array3(){
	sum := 0.0
	var avg float64
	xs := [6]float64{1, 2, 3, 4, 5, 6}
	switch len(xs) {
	case 0:
		avg = 0
	default:
		// 果你在遍历数组元素的时候，如果想遗弃索引id，可以直接把索引id标为下划线_
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	fmt.Println(avg)

	fmt.Println()
}

func main(){
	test_array1()
	test_array2()
	test_array3()
}
