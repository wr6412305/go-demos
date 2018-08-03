package return_multi_values

import "fmt"

func calc(a int, b int) (int, int){
	sum := a + b
	avg := sum / 2
	return sum, avg
}

func Return_multi_values(){
	sum, avg := calc(100, 200)
	fmt.Println("sum and avg is:", sum, avg)
	
	sum1, _ := calc(10, 20)
	fmt.Println("sum1:", sum1)
}