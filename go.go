package main

import "fmt"

func main() {
	var i, total int

	total = 0
	K := 6

	for i = 1; i <= K; i++ {
		if i%2 == 0 {
			total = total + (i * 2)
		} else {
			total = total - i

		}
	}
	fmt.Println("Total =", total)
}
