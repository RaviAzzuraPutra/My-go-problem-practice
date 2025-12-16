package main

import (
	"fmt"
	"strings"
)

func CountWords(sentence string) int {
	words := strings.Fields(sentence)

	return len(words)
}

func main() {
	kalimat := "CountWords menghitung jumlah kata dalam sebuah kalimat."

	fmt.Println(CountWords(kalimat))
}
