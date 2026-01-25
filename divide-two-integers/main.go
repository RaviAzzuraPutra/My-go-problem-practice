package main

import "fmt"

func Devide(dividend int, divisor int) int {

	result := float64(dividend) / float64(divisor)

	return int(result)

}

func main() {

	dividend := -90
	divisor := 8

	fmt.Println(Devide(dividend, divisor))

}
