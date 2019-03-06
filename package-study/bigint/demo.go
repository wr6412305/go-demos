package main

import (
	"fmt"
	"math"
	"math/big"
)

func demo1() {
	big1 := new(big.Int).SetUint64(uint64(1000))
	fmt.Println("big1 is: ", big1)

	big2 := big1.Uint64()
	fmt.Println("big2 is: ", big2)

	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	iq := big.NewInt(10000)
	ip.Mod(ip, iq)
	fmt.Printf("Big Int: %v\n", ip)
}
