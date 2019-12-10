package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// 随机数的三个阶段（随机安全性一阶更比一阶强）
// 随机性－－－不存在统计学偏差，是完全杂乱的数列
// 不可预测性－－－不能从过去的数列推测出下一个出现的数
// 不可重现性－－－除非将数列本身保存下来，否则不能重现相同的数列

// GetRand ...
func GetRand() {
	var seed int64 = 123456
	rand.Seed(seed)
	num1 := rand.Int()
	rand.Seed(time.Now().UnixNano())
	num2 := rand.Int() //[0,100)

	num3, err := crand.Int(crand.Reader, big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("仅满足随机数三大性质之一的随机性：")
	fmt.Println(num1)
	fmt.Println("-----------------------------------------")
	fmt.Println("满足随机数三大性质之一的随机性与不可预测性：")
	fmt.Println("虽然加入了当前时间作为种子，但是不可预测性非常弱")
	fmt.Println(num2)
	fmt.Println("go官方提供的更安全的随机数生成器,不可预测性较强，建议日常开发使用")
	fmt.Println(num3)
}
