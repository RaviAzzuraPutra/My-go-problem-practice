package main

import (
	"fmt"
	"strconv"
)

func StringCompression(s string) string {
	if len(s) == 0 {
		return s
	}

	var result string
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += string(s[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	result += string(s[len(s)-1]) + strconv.Itoa(count)

	if len(result) >= len(s) {
		return s
	}

	return result
}

func main() {
	Compression := "aaaaaaaaaaaaaaabbbccwwqqqqxbbbaa"
	Compression2 := "aaabbbcccaaa"
	Compression3 := "qaa"

	fmt.Println(StringCompression(Compression))
	fmt.Println(StringCompression(Compression2))
	fmt.Println(StringCompression(Compression3))
}
