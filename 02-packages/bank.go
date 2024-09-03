package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILENAME = "balance.txt"

func writeBalanceToFile(balance float64) {
	// Write the balance to a file
	balanceText := fmt.Sprint(balance)
	os.WriteFile(ACCOUNT_BALANCE_FILENAME, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(ACCOUNT_BALANCE_FILENAME)

	if err != nil {
		return 0, errors.New("error reading balance from file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("error parsing balance from file")
	}

	return balance, nil
}

func main() {
	var choice int
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		// panic(err)
	}

	fmt.Println("Welcome to Go Bank!")

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

			writeBalanceToFile(accountBalance)
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

			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}
