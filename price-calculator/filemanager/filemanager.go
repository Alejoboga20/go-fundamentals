package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	var lines []string

	file, err := os.Open(fm.InputFilePath)

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

func (fm *FileManager) WriteResult(data any) (any, error) {
	file, err := os.Create(fm.OutputFilePath)

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

func NewFileManger(inputFilePath string, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}
