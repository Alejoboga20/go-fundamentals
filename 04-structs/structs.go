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

type UserWithMethods struct {
	firstName string
	lastName  string
}

func (user UserWithMethods) outputUserDetailsOfUser() {
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// we can omit the field names if we know the order of the fields
	// appUser := User{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }
	appUser := User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}
	userWithMethod := UserWithMethods{
		firstName: userFirstName,
		lastName:  userLastName,
	}

	userWithMethod.outputUserDetailsOfUser()
	outputUserDetails(appUser)
	outputUserWithPointer(&appUser)
}

func outputUserDetails(appUser User) {
	fmt.Println("First Name:", appUser.firstName)
	fmt.Println("Last Name:", appUser.lastName)
	fmt.Println("Birthdate:", appUser.birthdate)
	fmt.Println("Created At:", appUser.createdAt)
}

func outputUserWithPointer(appUser *User) {
	// when using pointer with structs there is no need to use (*appUser).fieldName
	fmt.Println("First Name With Pointer:", appUser.firstName)
	fmt.Println("Last Name With Pointer:", appUser.lastName)
	fmt.Println("Birthdate With Pointer:", appUser.birthdate)
	fmt.Println("Created At With Pointer:", appUser.createdAt)
}

func getUserData(promptText string) string {
	var value string

	fmt.Print(promptText)
	fmt.Scan(&value)

	return value
}
