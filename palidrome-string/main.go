package main

import (
	"fmt"
	"strings"
)

func palindrome(text string) string {
	processWord := strings.ToLower(text)

	processWord = strings.ReplaceAll(processWord, " ", "")

	i := 0
	j := len(processWord) - 1

	for i < j {
		if processWord[i] != processWord[j] {
			return "not palindrome"
		}
		i++
		j--
	}

	return "palindrome"
}

func main() {
	fmt.Println("TEST", palindrome("TEST"))
	fmt.Println("katak", palindrome("katak"))
	fmt.Println("Sudah case insensitive", palindrome("Sudah case insensitive"))
	fmt.Println("Level", palindrome("Level"))
}
