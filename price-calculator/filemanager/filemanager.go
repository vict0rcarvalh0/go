package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLinesFromFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to encode data")
	}

	file.Close()
	return nil
}

func NewFileManager(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}