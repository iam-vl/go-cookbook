package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ExportLibDataForAnalysis(fn string, books []Book) error {
	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write header
	err = writer.Write([]string{"Title", "Author", "Pages"})
	if err != nil {
		return err
	}
	// Write book data
	for _, book := range books {
		err = writer.Write([]string{book.Title, book.Author, strconv.Itoa(book.Pages)})
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateLibSummary(books []Book) {
	fmt.Printf("Total books: %d\n", len(books))
	var totalPages int
	booksByAuthor := make(map[string][]Book)
	for _, book := range books {
		totalPages += book.Pages
		booksByAuthor[book.Author] = append(booksByAuthor[book.Author], book)
	}
	averagePages := float64(totalPages) / float64(len(books))
	fmt.Printf("Avg pages per book: %.2f\n", averagePages)
	for author, books := range booksByAuthor {
		fmt.Printf("%s has %d books.\n", author, len(books))
	}
}
