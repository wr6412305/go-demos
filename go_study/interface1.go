package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expense per month $%d\n", expense)
}

func interface2() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}

// 可以把接口想象成一个元组(type, value),type:接口包含的具体类型，value:接口包含的具体的值
type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func interface3() {
	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
}

// 任何类都实现了空接口interface{}
func describe1(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func interface4() {
	s := "hello"
	describe1(s)
	i := 55
	describe1(i)
	strt := struct {
		name string
	}{
		name: "ljs",
	}
	describe1(strt)
}
