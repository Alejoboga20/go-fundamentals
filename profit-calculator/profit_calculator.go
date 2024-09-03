package main

import (
	"fmt"
	"os"
	"strconv"
)

const FILE_NAME = "results.txt"

func main() {
	revenue := getUserInput("Enter the revenue: ")
	expenses := getUserInput("Enter the expenses: ")
	taxRate := getUserInput("Enter the tax rate: ")

	earnings_before_tax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax: ", earnings_before_tax)
	// print with formatting
	fmt.Printf("Profit: %v\nRatio: %v", profit, ratio)
	// We have different formatting options for printing values. We can use `%v` to print the value of the variable or `%f` to print the value of the variable as a float
	// we can also control the number of decimal places by using `%.nf` where `n` is the number of decimal places e.g. `%.2f` will print the value with 2 decimal places
	fmt.Printf("Profit: %f\nRatio: %.0f", profit, ratio)
	fmt.Println()
	// we can store the formatted string in a variable using `fmt.Sprintf`
	formatted_value := fmt.Sprintf("Profit: %f Ratio: %.0f", profit, ratio)
	fmt.Println("Formatted value: ", formatted_value)
	printText("Formatted value: ", formatted_value)
	fmt.Println()

	writeResultsInFile(revenue, expenses, taxRate)

	// We can also create multiline strings using backticks
	multiline_string := `
	This is a multiline string
	We can use it to write multiple lines of text
	without having to use the newline character`

	fmt.Println(multiline_string)
}

func printText(
	label string,
	text string,
) {
	fmt.Println(label, text)
}

func getUserInput(infoText string) float64 {
	var userInput string

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	validatedInput, _ := strconv.ParseFloat(userInput, 64)

	if validatedInput <= 0 {
		fmt.Println("Invalid input. Please enter a positive number")
		return getUserInput(infoText)
	}

	return validatedInput
}

func calculateFinancials(
	revenue,
	expenses,
	taxRate float64,
) (float64, float64, float64) {
	earnings_before_tax := revenue - expenses
	profit := earnings_before_tax * (1 - taxRate/100)
	ratio := earnings_before_tax / profit

	return earnings_before_tax, profit, ratio
}

func writeResultsInFile(revenue, expenses, taxRate float64) {
	fileInformation, _ := os.Stat(FILE_NAME)

	if fileInformation == nil {
		_, err := os.Create(FILE_NAME)

		if err != nil {
			fmt.Println("Error creating file: ", err)
		}
	}

	file, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("Revenue: %f\nExpenses: %f\nTax Rate: %f\n", revenue, expenses, taxRate))
}
