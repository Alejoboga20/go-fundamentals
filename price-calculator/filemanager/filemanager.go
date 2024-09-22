package filemanager

import (
	"bufio"
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
