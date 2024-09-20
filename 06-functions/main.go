package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 2, 3, 4, 5)
	anotherSum := sumup(1, numbers...)

	fmt.Println("Sum of numbers is", sum)
	fmt.Println("Another sum of numbers is", anotherSum)
}

func sumup(srartingValue int, numbers ...int) int {
	sum := 0

	fmt.Println("Starting value is", srartingValue)

	for _, number := range numbers {
		sum += number
	}

	return sum
}
