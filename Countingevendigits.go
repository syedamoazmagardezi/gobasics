package main

import "fmt"

func CountEvenDigits(number int) int {

	if number == 0 {
		return 0
	}

	lastDigit := number % 10

	if lastDigit%2 == 0 {
		return 1 + CountEvenDigits(number/10)
	} else {
		return CountEvenDigits(number / 10)
	}
}

func main() {

	fmt.Println(CountEvenDigits(48271))
	fmt.Println(CountEvenDigits(1359))
}
