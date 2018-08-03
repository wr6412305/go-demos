package main

import "fmt"

// 切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。
// 在内存管理方面，这是需要注意的。让我们假设我们有一个非常大的数组，
// 我们只想处理它的一小部分。然后，我们由这个数组创建一个切片，并
// 开始处理切片。这里需要重点注意的是，在切片引用时数组仍然存在内存中。
//一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个
// 切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收

func countries() []string {
	countries := []string {"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries) - 2]
	countriesCopy := make([]string, len(neededCountries))
	copy(countriesCopy, neededCountries)	// copies neededCountries to countriesCopy
	return countriesCopy
}

func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
