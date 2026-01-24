package main

import (
	"fmt"
	"strings"
)

func Palindrome(word string) string {
	lower := strings.ToLower(word)

	lower = strings.ReplaceAll(lower, " ", "")

	i := 0
	j := len(lower) - 1

	for i < j {
		if lower[i] != lower[j] {
			return "Is Not Palindrome!!"
		}
		i++
		j--
	}

	return "Is Palindrome!!"
}

func main() {
	Word1 := "CHELSEA JUARA"
	Word2 := "LEVEL"
	Word3 := "BACKSPACE"
	Word4 := "KATAK"
	Word5 := "Kasur Rusak"

	fmt.Println("CHELSEA : " + Palindrome(Word1))
	fmt.Println("LEVEL : " + Palindrome(Word2))
	fmt.Println("BACKSPACE : " + Palindrome(Word3))
	fmt.Println("KATAK : " + Palindrome(Word4))
	fmt.Println("Kasur Rusak : " + Palindrome(Word5))
}
