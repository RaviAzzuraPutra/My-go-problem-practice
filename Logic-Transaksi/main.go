package main

import "fmt"

func Transaction(balance int, price int) string {

	if price < 0 && balance < 0 {
		return "Invalid Input"

	}

	if balance < price {
		return "Insufficient Balance"
	}

	if balance >= price {
		return "Success"
	}

	return ""

}

func main() {

	balance1 := 100000
	price1 := 2000

	balance2 := 1
	price2 := 2000

	balance3 := -100000
	price3 := -2000

	fmt.Println(Transaction(balance1, price1))
	fmt.Println(Transaction(balance2, price2))
	fmt.Println(Transaction(balance3, price3))

}
