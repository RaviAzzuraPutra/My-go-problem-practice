package main

import "fmt"

func SumArray(nums []int) int {

	var result int

	for _, value := range nums {
		result = result + value
	}

	return result

}

func main() {
	array := []int{1, 2, 3, 4, 5, 99}

	fmt.Println(SumArray(array))
}
