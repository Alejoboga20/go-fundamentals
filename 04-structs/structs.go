package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(appUser)
}

func outputUserDetails(appUser User) {
	fmt.Println("First Name:", appUser.firstName)
	fmt.Println("Last Name:", appUser.lastName)
	fmt.Println("Birthdate:", appUser.birthdate)
	fmt.Println("Created At:", appUser.createdAt)
}

func getUserData(promptText string) string {
	var value string

	fmt.Print(promptText)
	fmt.Scan(&value)

	return value
}
