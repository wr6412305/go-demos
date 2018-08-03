package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["India"] = "新德里"
	countryCapitalMap["China"] = "北京"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 查看元素在集合中是否存在
	captial, ok := countryCapitalMap["美国"]
	if ok {
		fmt.Println("美国的首都是", captial)
	}else {
		fmt.Println("map中未存储美国的首都")
	}
}
