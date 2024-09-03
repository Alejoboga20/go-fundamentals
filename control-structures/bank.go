package main

import (
	"fmt"
)

func main() {
	var choice int
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What woud you like to do?")

	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// wantsToCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is $", accountBalance)
	}
}
