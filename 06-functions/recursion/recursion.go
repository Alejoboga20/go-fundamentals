package recursion

import "fmt"

func main() {
	fmt.Println("Factorial of 5 is", factorial(5))
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// 5 * factorial(4)
// 5 * 4 * factorial(3)
// 5 * 4 * 3 * factorial(2)
// 5 * 4 * 3 * 2 * factorial(1)
// 5 * 4 * 3 * 2 * 1 * factorial(0)
// 5 * 4 * 3 * 2 * 1 * 1
