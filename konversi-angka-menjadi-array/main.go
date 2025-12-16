package main

import (
	"fmt"
	"strconv"
)

func convertInteggerToArray(nums int) []int {
	IntToStr := strconv.Itoa(nums)

	result := []int{}

	for _, value := range IntToStr {
		StrToInt, _ := strconv.Atoi(string(value))

		result = append(result, StrToInt)
	}

	return result
}

func main() {
	Numbers := 1234567890

	fmt.Println(convertInteggerToArray(Numbers))
}
