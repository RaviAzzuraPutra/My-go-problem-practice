package main

import (
	"fmt"
	"sort"
)

func SortingProductID(nums []int, k int) []int {

	product := make(map[int]int)

	for _, value := range nums {
		product[value]++
	}

	type kv struct {
		ID   int
		Freq int
	}

	var ss []kv

	for id, freq := range product {
		ss = append(ss, kv{id, freq})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Freq > ss[j].Freq
	})

	var result []int

	for i := 0; i < k && i < len(ss); i++ {
		result = append(result, ss[i].ID)
	}

	return result

}

func main() {

	array := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 10}
	k := 5

	fmt.Println(SortingProductID(array, k))

}
