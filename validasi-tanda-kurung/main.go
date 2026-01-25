package main

import "fmt"

func Validation(s string) bool {
	stack := []rune{}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if opening, found := pairs[char]; found {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {

	input := "({[]})"
	input2 := "([})"

	fmt.Println(Validation(input))
	fmt.Println(Validation(input2))

}
