package main

import "fmt"

func main() {
	prices := []float64{1.50, 2.35, 3.123, 4, 5}

	for i := range prices {
		fmt.Println(prices[i])
	}
}
