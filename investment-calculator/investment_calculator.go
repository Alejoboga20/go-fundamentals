package main

import (
	"fmt"
	"math"
)

func main() {
	// We can use `const` to define constants. Constants are immutable and their value cannot be changed
	const inflationRate = 2.5
	// We can define more than one variable in the same line. We can use the same type of inference for multiple variables
	var investmentAmount float64
	years := 5.0

	fmt.Print("Enter the investment amount: ")
	// We need to use `&` to get the memory address of the variable. Scan will read the input from the user and store it in the memory address of the variable
	fmt.Scan(&investmentAmount)

	// This `:=` is used for variables where the type is not explicitly defined but should be inferred from the value assigned to it
	expectedReturnRate := 5.5
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Invested amount: ", investmentAmount)
	fmt.Println("Future value is: ", futureValue)
	fmt.Println("Future real value is: ", futureRealValue)
}
