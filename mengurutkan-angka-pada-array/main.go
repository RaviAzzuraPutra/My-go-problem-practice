package main

import (
	"fmt"
	"sort"
)

func SortArray(nums []int) []int {
	sort.Ints(nums)

	return nums
}

func main() {
	array := []int{3, 1, 5, 2, 4, 99, 0}

	fmt.Println(SortArray(array[:]))
}
