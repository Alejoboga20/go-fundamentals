package main

import (
	"fmt"

	"example.com/bank/fileutils"
	"github.com/Pallinder/go-randomdata"
)

const ACCOUNT_BALANCE_FILENAME = "balance.txt"

func main() {
	var choice int
	var accountBalance, err = fileutils.GetFloatFromFile(ACCOUNT_BALANCE_FILENAME)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		// panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Our phone number is:", randomdata.PhoneNumber())

	for {
		presentOptions()
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

			fileutils.WriteFloatToFile(ACCOUNT_BALANCE_FILENAME, accountBalance)
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

			fileutils.WriteFloatToFile(ACCOUNT_BALANCE_FILENAME, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}
