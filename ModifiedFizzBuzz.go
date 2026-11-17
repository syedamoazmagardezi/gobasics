package main

import "fmt"

func main() {

	var N int

	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("COMPUTATION")

		} else if i%3 == 0 {
			fmt.Println("ALGORITHM")

		} else if i%5 == 0 {
			fmt.Println("LOGIC")

		} else {
			fmt.Println(i)
		}
	}
}
