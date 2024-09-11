package main

import (
	"bufio"
	"fmt"
	"notes-app/note"
	"notes-app/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	todoText := getTodoData()
	title, content := getNoteData()

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(&userTodo)
	outputData(&userNote)

	printSomething("Note and Todo saved successfully!")
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo text:")

	return text
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Print("Data saved successfully!")

	return nil
}

func outputData(data outputtable) {
	data.Display()
	saveData(data)
}

func printSomething(data interface{}) {
	typedValue, ok := data.(int) // extracting the value and checking the type

	if ok {
		fmt.Println("Integer value:", typedValue)
		return
	}

	switch data.(type) {
	case string:
		fmt.Println("String value")
	case int:
		fmt.Println("Integer value")
	case bool:
		fmt.Println("Boolean value")
	default:
		fmt.Println("Unknown type")
	}

	fmt.Println(data)

	result := add(1, 2)
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
