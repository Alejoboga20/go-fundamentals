package main

import "fmt"

func presentOptions() {
	fmt.Println("What woud you like to do?")

	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	fmt.Print("Your choice: ")
}
