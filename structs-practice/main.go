package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Errorf("Error getting note data: %v", err)
	}

}

func getUserInput(prompt string) (string, error) {
	var input string

	fmt.Println(prompt)
	fmt.Scanln(&input)

	if input == "" {
		return "", errors.New("Input cannot be empty")

	}

	return input, nil
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:")

	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note content:")

	if err != nil {
		return "", "", err
	}

	return title, content, nil
}
