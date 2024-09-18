# Collections

In Go, collections are data structures that allow you to store multiple values in a single variable. The most common collections in Go are arrays, slices, and maps.

## Arrays

An array is a fixed-size collection of elements of the same type. The size of an array is part of its type, so arrays cannot be resized. Here's how you can declare an array in Go:

```go
var numbers [5]int // we can declare an array of 5 integers
prices := []float64{1.50, 2.35, 3.123, 4, 5} // we can also declare an array using the shorthand syntax

for i := range prices {
		fmt.Println(prices[i])
}
```

### Slices

A Slice is a window in the original array. It is a reference to the original array, so any changes made to the slice will affect the original array. Here's how you can declare a slice in Go:

```go
featuredPricesOne := prices[1:3]
featuredPricesTwo := prices[:3]
featuredPricesThree := prices[1:]
concatenatedPrices := featuredPricesOne[:1]
fmt.Println(featuredPricesOne, featuredPricesTwo, featuredPricesThree, concatenatedPrices)
```

## Maps

Maps in Go are unordered collections of key-value pairs. Here's how you can declare a map in Go:

```go
package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Azure":    "https://www.azure.com",
	}
	mapWithNumbersAsKeys := map[int]string{
		1: "One",
		2: "Two",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])
	fmt.Println(mapWithNumbersAsKeys)
	fmt.Println(mapWithNumbersAsKeys[1])

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Facebook")
	fmt.Println(websites)
}
```

Different from arrays and slices, maps are reference types. When you assign a map to another variable or pass it to a function, you are passing a reference to the original map. This means that any changes made to the map will affect the original map. Also maps are dynamic, you can add or remove elements from a map after it has been created.
