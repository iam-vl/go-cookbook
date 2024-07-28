package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func SaveBooks(fn string, books []Book) error {
	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, book := range books {
		jsonData, err := json.Marshal(book)
		if err != nil {
			return err
		}
		_, err = writer.WriteString(string(jsonData) + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
func LoadBooks(fn string) ([]Book, error) {
	var books []Book
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var book Book
		if err := json.Unmarshal([]byte(scanner.Text()), &book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, scanner.Err()
}
