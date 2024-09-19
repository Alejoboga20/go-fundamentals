package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// factory function
	doubler := createTransformer(2)
	tripler := createTransformer(3)

	doubled := transformNumbers(&numbers, doubler)
	tripled := transformNumbers(&numbers, tripler)

	// anonymous function
	transformed := transformNumbers(&numbers, func(value int) int {
		return value * 2
	})

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := []int{}

	for _, val := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(val))
	}

	return transformedNumbers
}

// factory function
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
