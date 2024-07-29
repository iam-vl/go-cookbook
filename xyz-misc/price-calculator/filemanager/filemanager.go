package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // *bufio.Scanner

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Failed to read line in file")
	}
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON.")
	}

	file.Close()
	return nil
}
