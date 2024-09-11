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
