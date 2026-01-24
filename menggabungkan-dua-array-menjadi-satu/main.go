package main

import (
	"fmt"
	"sort"
)

func MergeArray(nums1, nums2 []int) []int {

	result := make([]int, len(nums1)+len(nums2))

	copy(result, nums1)

	copy(result[len(nums1):], nums2)

	sort.Ints(result)

	return result
}

func main() {
	array1 := []int{1, 99, 34, 7, 2}
	array2 := []int{2, 0, 91, 23, 67}

	fmt.Println(MergeArray(array1, array2))
}
