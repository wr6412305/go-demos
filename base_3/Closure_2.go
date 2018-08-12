package main

import (
	"fmt"
	)

type student struct {
	firstname string
	lastname string
	grade string
	country string
}

func filter(s [] student, f func(student) bool) [] student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []int, f func(int) int) [] int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	s1 := student{
		firstname: "Naveen",
		lastname:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstname: "Samuel",
		lastname:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}

	f1 := func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	}

	s_res := filter(s, f1)
	fmt.Println(s_res)

	a := []int { 5, 6, 7, 8, 9 }
	fMulti5 := func(a int) int {
		return a * 5
	}
	r := iMap(a, fMulti5)
	fmt.Println(r)
}
