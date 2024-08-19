package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	earnings__before_tax := revenue - expenses
	profit := earnings__before_tax * (1 - taxRate/100)
	ratio := earnings__before_tax / profit

	fmt.Println("Earnings before tax: ", earnings__before_tax)
	// print with formatting
	fmt.Printf("Profit: %v\nRatio: %v", profit, ratio)
}
