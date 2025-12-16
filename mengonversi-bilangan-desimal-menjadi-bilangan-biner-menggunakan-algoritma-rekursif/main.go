package main

import (
	"fmt"
	"strconv"
)

func ConvertionDecimalToBinary(nums int) string {
	if nums == 0 {
		return ""
	}

	var result string

	result = ConvertionDecimalToBinary(nums/2) + strconv.Itoa(nums%2)

	return result
}

func main() {
	decimal := 222

	fmt.Println(ConvertionDecimalToBinary(decimal))
}
