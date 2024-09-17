package main

import "fmt"

func main() {
	productNames := [4]string{"laptop", "mouse", "keyboard"}
	prices := []float64{1.50, 2.35, 3.123, 4, 5}

	fmt.Println(productNames)
	productNames[0] = "monitor"
	fmt.Println(productNames)

	for i := range prices {
		fmt.Println(prices[i])
	}

	// Slices
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
