package main

import (
	"bufio"
	"os"
	"regexp"
)

var bookDetailsPattern = regexp.MustCompile(`Title: (.+), Author: (.+), Pages: (\d+)`)

func OpenFileScan(filename string) (*os.File, *bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	return file, scanner, nil
}
