package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILENAME = "balance.txt"

func writeFloatToFile(fileName string, floatValue float64) {
	dataText := fmt.Sprint(floatValue)
	os.WriteFile(fileName, []byte(dataText), 0644)
}

func getFloatFromFile(fileName string, defaultValue ...float64) (float64, error) {
	// ...float64 is a variadic parameter: it allows the function to accept 0 or more float64 values
	var defVal float64
	data, err := os.ReadFile(fileName)

	if len(defaultValue) > 0 {
		defVal = defaultValue[0]
	} else {
		defVal = 0
	}

	if err != nil {
		return defVal, errors.New("error reading from file")
	}

	dataString := string(data)
	dataFloat, err := strconv.ParseFloat(dataString, 64)

	if err != nil {
		return defVal, errors.New("error parsing from file")
	}

	return dataFloat, nil
}

func main() {
	var choice int
	var accountBalance, err = getFloatFromFile(ACCOUNT_BALANCE_FILENAME)

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

			writeFloatToFile(ACCOUNT_BALANCE_FILENAME, accountBalance)
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

			writeFloatToFile(ACCOUNT_BALANCE_FILENAME, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}
