package main

import "fmt"

func main() {
	sum := sumup(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers is", sum)
}

func sumup(srartingValue int, numbers ...int) int {
	sum := 0

	fmt.Println("Starting value is", srartingValue)

	for _, number := range numbers {
		sum += number
	}

	return sum
}
