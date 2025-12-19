package main

import "fmt"

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FIZZ-BUZZ")
		} else if i%5 == 0 {
			fmt.Println("BUZZ")
		} else if i%3 == 0 {
			fmt.Println("FIZZ")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	FizzBuzz()
}
