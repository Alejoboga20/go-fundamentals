package main

import "fmt"

// func main() {
// 	productNames := [4]string{"laptop", "mouse", "keyboard"}
// 	prices := []float64{1.50, 2.35, 3.123, 4, 5}

// 	fmt.Println(productNames)
// 	productNames[0] = "monitor"
// 	fmt.Println(productNames)

// 	for i := range prices {
// 		fmt.Println(prices[i])
// 	}

// 	// Slices
// 	featuredPricesOne := prices[1:3]
// 	featuredPricesTwo := prices[:3]
// 	featuredPricesThree := prices[1:]
// 	concatenatedPrices := featuredPricesOne[:1]

// 	fmt.Println(featuredPricesOne, featuredPricesTwo, featuredPricesThree, concatenatedPrices)

// 	// Slices metadata
// 	fmt.Println(len(prices), cap(prices))
// 	fmt.Println(len(featuredPricesOne), cap(featuredPricesOne))

// 	// we can re-slice and select more elements to the right
// 	concatenatedPrices = concatenatedPrices[:3]
// 	fmt.Println(featuredPricesOne, concatenatedPrices)

// }

func main() {
	// if we don't specify the size of the array, it will be a slice, not an array
	// with an slice we can add or remove elements, no matter the size
	prices := []float64{10.99, 5.99, 3.99}
	fmt.Println(prices[0])

	// we can not add elements to indexes that are not in the slice e.g. prices[3] = 1.99
	// to add elements we use the append function
	updatedPrices := append(prices, 1.99)
	fmt.Println(updatedPrices, prices)
}
