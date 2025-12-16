package main

import "fmt"

func BigArrayNumber(nums []int) int {
	var result int

	for i := 0; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}

func main() {
	array := []int{1, 301, 5, 3, 7, 911, 210, 666}

	fmt.Println(BigArrayNumber(array[:]))
}
