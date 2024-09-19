package functionsarevalues

import "fmt"

type transformFunc func(int) int

func main() {
	numbersToDouble := []int{1, 2, 3, 4, 5}
	numbersToTriple := []int{3, 4, 5, 6, 7}

	doubled := transformNumbers(&numbersToDouble, getTransformerFunction(&numbersToDouble))
	tripled := transformNumbers(&numbersToTriple, getTransformerFunction(&numbersToTriple))

	fmt.Println("Original numbers to double:", numbersToDouble)
	fmt.Println("Original numbers to triple:", numbersToTriple)
	fmt.Println("Doubled numbers:", doubled)
	fmt.Println("Tripled numbers:", tripled)
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	transformedNumbers := []int{}

	for _, value := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(value))
	}

	return transformedNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return double
	}

	return triple
}
