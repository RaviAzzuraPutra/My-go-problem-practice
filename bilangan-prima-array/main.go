package main

import (
	"fmt"
	"math"
)

func is_prime(n int) bool {
	if n < 2 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func filter_prime_numbers(numbers []int) []int {
	primes := []int{}

	for _, value := range numbers {
		if is_prime(value) {
			primes = append(primes, value)
		}
	}

	return primes
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{11, 12, 13, 14, 15, 16, 17}
	arr3 := []int{4, 6, 8, 10, 12}

	fmt.Println(filter_prime_numbers(arr1)) // Output: [2 3 5 7]
	fmt.Println(filter_prime_numbers(arr2)) // Output: [11 13 17]
	fmt.Println(filter_prime_numbers(arr3)) // Output: []
}
