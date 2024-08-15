package main

import (
	"fmt"
	"math"
)

func main() {
	// We can use `const` to define constants. Constants are immutable and their value cannot be changed
	const inflationRate = 2.5
	// We can define more than one variable in the same line. We can use the same type of inference for multiple variables
	var investmentAmount, years float64
	// we can init a variable with a value using `:=` for inference
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	// We need to use `&` to get the memory address of the variable. Scan will read the input from the user and store it in the memory address of the variable
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years you'd like to invest: ")
	fmt.Scan((&years))

	// Note: `fmt.Scan` only allows us to read one input at a time or simple words
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan((&expectedReturnRate))

	// This `:=` is used for variables where the type is not explicitly defined but should be inferred from the value assigned to it
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Invested amount: ", investmentAmount)
	fmt.Println("Future value is: ", futureValue)
	fmt.Println("Future real value is: ", futureRealValue)
}
