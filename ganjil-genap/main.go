package main

import "fmt"

func OddAndEven(nums []int) ([]int, []int) {

	var OddNumber []int
	var EvenNumber []int

	for _, value := range nums {
		if value%2 == 0 {
			EvenNumber = append(EvenNumber, value)
		} else {
			OddNumber = append(OddNumber, value)
		}
	}

	return OddNumber, EvenNumber

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Println(OddAndEven(array))
}
