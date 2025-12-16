package main

import "fmt"

func reverse_string(text string) string {
	var reverse string

	for _, v := range text {
		reverse = string(v) + reverse
	}

	return reverse
}

func main() {
	str := "LION STAR"

	fmt.Println(reverse_string(str))
}
