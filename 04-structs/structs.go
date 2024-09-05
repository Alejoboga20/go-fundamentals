package main

import (
	"errors"
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

func (user *UserWithMethods) clearUserName() {
	user.firstName = ""
	user.lastName = ""
}

// constructor function: it's a pattern to create a new instance of a struct
func newUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		fmt.Println("Please provide all the required fields")

		return nil, errors.New("missing required fields")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println("Error creating user")
		return
	}

	userWithMethod := UserWithMethods{
		firstName: userFirstName,
		lastName:  userLastName,
	}

	userWithMethod.outputUserDetailsOfUser()
	userWithMethod.clearUserName()
	userWithMethod.outputUserDetailsOfUser()

	outputUserDetails(*appUser)
	outputUserWithPointer(appUser)

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
	fmt.Scanln(&value)

	return value
}
