package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(value int) int {
		return value * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := []int{}

	for _, val := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(val))
	}

	return transformedNumbers
}
