package main

import (
	"fmt"
	"strings"
)

func MostLetter(word string) string {
	lower := strings.ToLower(word)

	MapOfLetter := make(map[rune]int)

	for _, value1 := range lower {
		MapOfLetter[value1]++
	}

	var maxCount int
	var MostLetterOfTheWord rune

	for key2, value2 := range MapOfLetter {
		if value2 > maxCount {
			maxCount = value2
			MostLetterOfTheWord = key2
		}
	}

	return string(MostLetterOfTheWord)
}

func main() {
	letter := "HELLO"

	fmt.Println(MostLetter(letter))
}
