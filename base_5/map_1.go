package main

import (
	"fmt"
	"sort"
)

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	countryCapitalMap["China"] = "Beijing"

	fmt.Println("Original:")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	fmt.Println()

	delete(countryCapitalMap, "India")
	fmt.Println("After deletion:")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	fmt.Println()

	capitalCountryMap := make(map[string]string, len(countryCapitalMap))
	for country, capital := range countryCapitalMap {
		capitalCountryMap[capital] = country
	}
	fmt.Println("After the reversal:")
	for capital := range capitalCountryMap {
		fmt.Println("Country of", capital, "is", capitalCountryMap[capital])
	}
	fmt.Println()

	// 查看元素在集合中是否存在
	capital, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("Capital fo United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}
	fmt.Println()

	// 排序输出
	var keys []string
	for k := range countryCapitalMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", countryCapitalMap[k])
	}
}
