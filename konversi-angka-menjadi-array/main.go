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

func convertArrayToInt(nums []int) int {

	var result string

	for _, value := range nums {
		result += strconv.Itoa(value)
	}

	StrToInt, _ := strconv.Atoi(result)

	return StrToInt

}

func main() {
	Numbers := 1234567890
	Array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(convertInteggerToArray(Numbers))
	fmt.Println(convertArrayToInt(Array[:]))
}
