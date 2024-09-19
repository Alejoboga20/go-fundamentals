package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("Original numbers:", numbers)
	fmt.Println("Doubled numbers:", doubled)
	fmt.Println("Tripled numbers:", tripled)
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	doubledNumbers := []int{}

	for _, value := range *numbers {
		doubledNumbers = append(doubledNumbers, transform(value))
	}

	return doubledNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
