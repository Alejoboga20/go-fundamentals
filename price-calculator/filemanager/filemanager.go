package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("an error occurred while opening the file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("an error occurred while reading the file")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data any) (any, error) {
	file, err := os.Create(path)

	if err != nil {
		return nil, errors.New("an error occurred while creating the file")
	}

	enconder := json.NewEncoder(file)
	err = enconder.Encode(data)

	if err != nil {
		file.Close()
		return nil, errors.New("failed to encode the data")
	}

	file.Close()
	return nil, nil
}
