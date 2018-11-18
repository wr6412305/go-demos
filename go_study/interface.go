package main

import (
	"fmt"
)

// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func interface1() {
	name := MyString("asdfahsdgfa")
	var v VowelsFinder
	v = name // MyString 实现了接口VowelsFinder, 所以可以将MyString转换为VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
}
