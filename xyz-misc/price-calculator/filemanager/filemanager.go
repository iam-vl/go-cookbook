package filemanager

import (
	"bufio"
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
	// func (s *bufio.Scanner) Scan() bool
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Failed to read line in file")
	}
	return lines, nil
}
