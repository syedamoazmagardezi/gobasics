package main

import "fmt"

func main() {

	var N int
	var temperature float64

	idealCount := 0
	overheatedCount := 0
	totalTemperature := 0.0

	fmt.Print("Enter total minutes: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {

		fmt.Print("Enter temperature: ")
		fmt.Scan(&temperature)

		totalTemperature = totalTemperature + temperature

		if temperature >= 195 && temperature <= 210 {
			idealCount++
		} else if temperature > 210 {
			overheatedCount++
		}
	}

	average := totalTemperature / float64(N)

	fmt.Println("Ideal Minutes =", idealCount)
	fmt.Println("Overheated Minutes =", overheatedCount)
	fmt.Println("Average Temperature =", average)
}
