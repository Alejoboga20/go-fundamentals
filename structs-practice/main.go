package main

import (
	"fmt"
	"notes-app/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUserInput(prompt string) string {
	var input string

	fmt.Println(prompt)
	fmt.Scanln(&input)

	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}
