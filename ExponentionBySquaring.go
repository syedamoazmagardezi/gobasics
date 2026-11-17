package main

import "fmt"

func FastPower(x int, n int) int {

	if n == 0 {
		return 1
	}

	if n%2 == 0 {

		half := FastPower(x, n/2)

		return half * half

	} else {

		return x * FastPower(x, n-1)
	}
}

func main() {

	result := FastPower(2, 6)

	fmt.Println("Final Answer =", result)
}
