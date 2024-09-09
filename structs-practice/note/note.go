package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note *Note) Display() {
	fmt.Printf("Your note title %v has the following content: \n\n %v", note.title, note.content)
}

func (note *Note) Save() {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + ".txt"

	os.WriteFile(fileName, []byte(note.content), 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("input cannot be empty")

	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
