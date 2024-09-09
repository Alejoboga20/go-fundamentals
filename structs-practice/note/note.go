package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note *Note) Display() {
	fmt.Printf("Your note title %v has the following content: \n\n %v \n", note.Title, note.Content)
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonNote, err := json.Marshal(note)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return os.WriteFile(fileName, jsonNote, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("input cannot be empty")

	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
