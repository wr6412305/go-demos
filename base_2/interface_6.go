package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type myString string

// myString implements VowelsFinder
func (ms myString) FindVowels() [] rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := myString("Sam Anderson")
	var v VowelsFinder
	// myString类型实现了接口VowelsFinder，所以这里是合法的
	v = name		// possible since myString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", v.FindVowels())
	fmt.Printf("Vowels are %c\n", name.FindVowels())
}
