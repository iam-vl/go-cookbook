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

func Get2Books() []Book {
	b1 := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. Donovan",
		Pages:  300,
	}
	b2 := Book{Title: "Go in Action", Author: "Bill Kennedy", Pages: 300}
	return []Book{b1, b2}
}
