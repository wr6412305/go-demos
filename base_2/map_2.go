package main

import "fmt"

// delete删除集合中的元素

func main() {
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}

	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}
}
