package main

import (
	"fmt"
)

func main() {
	var choice int
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What woud you like to do?")

		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is $", accountBalance)
		case 2:
			var depositAmount float64

			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("You cannot deposit a negative amount")
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Your new balance is $", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Println("Enter the amount you want to withdraw: ")

			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("You cannot withdraw a negative amount")
				continue
			} else if withdrawAmount > accountBalance {
				fmt.Println("You do not have enough balance to withdraw that amount")
				continue
			}

			accountBalance -= withdrawAmount // accountBalance = accountBalance - withdrawAmount
			fmt.Println("Your new balance is $", accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}

}

// func main() {
// 	var choice int
// 	var accountBalance = 1000.0

// 	fmt.Println("Welcome to Go Bank!")

// 	for {
// 		fmt.Println("What woud you like to do?")

// 		fmt.Println("1. Check Balance")
// 		fmt.Println("2. Deposit Money")
// 		fmt.Println("3. Withdraw Money")
// 		fmt.Println("4. Exit")

// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		// wantsToCheckBalance := choice == 1

// 		if choice == 1 {
// 			fmt.Println("Your balance is $", accountBalance)
// 		} else if choice == 2 {
// 			var depositAmount float64

// 			fmt.Print("Enter the amount you want to deposit: ")
// 			fmt.Scan(&depositAmount)

// 			if depositAmount <= 0 {
// 				fmt.Println("You cannot deposit a negative amount")
// 				continue
// 			}

// 			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
// 			fmt.Println("Your new balance is $", accountBalance)
// 		} else if choice == 3 {
// 			var withdrawAmount float64
// 			fmt.Println("Enter the amount you want to withdraw: ")

// 			fmt.Scan(&withdrawAmount)

// 			if withdrawAmount <= 0 {
// 				fmt.Println("You cannot withdraw a negative amount")
// 				continue
// 			} else if withdrawAmount > accountBalance {
// 				fmt.Println("You do not have enough balance to withdraw that amount")
// 				continue
// 			}

// 			accountBalance -= withdrawAmount // accountBalance = accountBalance - withdrawAmount
// 			fmt.Println("Your new balance is $", accountBalance)
// 		} else if choice == 4 {
// 			fmt.Println("Goodbye!")
// 			break
// 		} else {
// 			fmt.Println("Invalid choice. Please try again.")
// 		}
// 	}

// 	fmt.Println("Thank you for using Go Bank!")
// }
