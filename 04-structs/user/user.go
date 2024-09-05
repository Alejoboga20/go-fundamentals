package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

type UserWithMethods struct {
	FirstName string
	LastName  string
}

func (user UserWithMethods) OutputUserDetailsOfUser() {
	fmt.Println("First Name:", user.FirstName)
	fmt.Println("Last Name:", user.LastName)
}

func (user *UserWithMethods) ClearUserName() {
	user.FirstName = ""
	user.LastName = ""
}

// constructor function: it's a pattern to create a new instance of a struct
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		fmt.Println("Please provide all the required fields")

		return nil, errors.New("missing required fields")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}
