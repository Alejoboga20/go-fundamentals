package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo *Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo *Todo) Save() error {
	fileName := "todo.json"

	jsonTodo, err := json.Marshal(todo)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return os.WriteFile(fileName, jsonTodo, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("input cannot be empty")

	}
	return Todo{
		Text: text,
	}, nil
}
