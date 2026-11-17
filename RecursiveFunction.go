package main

import "fmt"

func ComputePattern(n int) int {

	if n <= 1 {
		return 1
	} else if n%2 == 1 {
		return n + ComputePattern(n-1)
	} else {
		return ComputePattern(n-2) * 2
	}
}

func main() {
	M := 3

	result := ComputePattern(M)

	fmt.Println("Final Answer =", result)
}
