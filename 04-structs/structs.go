package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println("Error creating user")
		return
	}

	userWithMethod := user.UserWithMethods{
		FirstName: userFirstName,
		LastName:  userLastName,
	}

	userWithMethod.OutputUserDetailsOfUser()
	userWithMethod.ClearUserName()
	userWithMethod.OutputUserDetailsOfUser()

	outputUserDetails(*appUser)
	outputUserWithPointer(appUser)

}

func outputUserDetails(appUser user.User) {
	fmt.Println("First Name:", appUser.FirstName)
	fmt.Println("Last Name:", appUser.LastName)
	fmt.Println("Birthdate:", appUser.Birthdate)
	fmt.Println("Created At:", appUser.CreatedAt)
}

func outputUserWithPointer(appUser *user.User) {
	// when using pointer with structs there is no need to use (*appUser).fieldName
	fmt.Println("First Name With Pointer:", appUser.FirstName)
	fmt.Println("Last Name With Pointer:", appUser.LastName)
	fmt.Println("Birthdate With Pointer:", appUser.Birthdate)
	fmt.Println("Created At With Pointer:", appUser.CreatedAt)
}

func getUserData(promptText string) string {
	var value string

	fmt.Print(promptText)
	fmt.Scanln(&value)

	return value
}
