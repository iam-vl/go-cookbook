package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

// For XML, we often use a wrapper type for a collection
type Library struct {
	Books []Book `xml:"book"`
}

func ParseBooksFromFile(fn string) ([]Book, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Printf("File Type: %T\n", file)
	fmt.Printf("Scanner Type: %T\n", scanner)
	var books []Book

	for scanner.Scan() {
		matches := bookDetailsPattern.FindStringSubmatch(scanner.Text())
		if matches != nil && len(matches) == 4 {
			title := matches[1]
			author := matches[2]
			pages, err := strconv.Atoi(matches[3])
			if err != nil {
				// Log error and continue parsing the rest of file
				fmt.Printf("Invalid page number for book '%s': %s\n", title, err)
				continue
			}
			books = append(books, Book{Title: title, Author: author, Pages: pages})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func ExportBooksToXml(books []Book) (string, error) {
	lib := Library{Books: books}
	xmlData, err := xml.MarshalIndent(lib, "", "")
	if err != nil {
		return "", err
	}
	return string(xmlData), nil
}

func ImportBooksFromXml(xmlData string) ([]Book, error) {
	var lib Library
	err := xml.Unmarshal([]byte(xmlData), &lib)
	if err != nil {
		return nil, err
	}
	return lib.Books, nil
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
