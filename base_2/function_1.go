package main

import (
	"math"
	"fmt"
)

func main(){
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}
