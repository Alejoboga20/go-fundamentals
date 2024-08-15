package main

import (
	"fmt"
	"math"
)

func main() {
	// We can define more than one variable in the same line. We can use the same type of inference for multiple variables
	var investmentAmount, years float64 = 1000, 10

	// This `:=` is used for variables where the type is not explicitly defined but should be inferred from the value assigned to it
	expectedReturnRate := 5.5
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println("Invested amount: ", investmentAmount)
	fmt.Println("Future value is: ", futureValue)
}
