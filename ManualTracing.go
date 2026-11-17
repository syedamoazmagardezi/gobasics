package main

import "fmt"

func main() {
	total := 0
	K := 10

	for i := 1; i <= K; i++ {

		if i%2 == 0 {
			total = total + (i * 2)
		} else {
			total = total - i
		}

		fmt.Println("i =", i, "Total =", total)
	}

	fmt.Println("Final Total =", total)
}
