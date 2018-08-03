package main

import "fmt"

// 数组是值类型，并且按值传递
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function", num)
}

func printarray(a [3][2] string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
}

func main() {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function", num)
	changeLocal(num)
	fmt.Println("after passing to function", num)

	sumArr := float64(0)
	for i, v := range num {	// range returns both the index and value
		fmt.Printf("%d the element of num is %.2f\n", i, float64(v))
		// Go语言没有隐式类型转换，必须显示强制转换
		sumArr += float64(v)
	}
	fmt.Println("sum of all elements of a", sumArr)

	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},	// 最后这个逗号是必须的
	}
	printarray(a)
	fmt.Println()

	var b[3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	printarray(b)
}
