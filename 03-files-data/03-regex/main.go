package main

import "fmt"

func main() {
	fn := "book_listings.txt"
	books, err := ParseBooksFromFile(fn)
	if err != nil {
		fmt.Println("Error parsing books from file:", err)
		return
	}
	for _, book := range books {
		fmt.Printf("Parsed book: %+v\n", book)
	}
}

func checkXml() {
	books := []Book{
		{"The Go Programming Language", "Alan A. Donovan", 380},
		{"Go in Action", "Bill Kennedy", 300},
	}
	xmlOutput, err := ExportBooksToXml(books)
	if err != nil {
		fmt.Println("Error exporting books to XML:", err)
		return
	}
	fmt.Println("XML output:", xmlOutput)

	// Simulate importing
	importedBooks, err := ImportBooksFromXml(xmlOutput)
	if err != nil {
		fmt.Println("Error importing books from XML:", err)
		return
	}
	fmt.Println("Imported books:", importedBooks)

}

func CheckSaveLoadBooks() {
	books := []Book{
		{"The Go Programming Language", "Alan A. Donovan", 380},
		{"Go in Action", "Bill Kennedy", 300},
	}
	fn := "books.json"
	// Save books to file
	err := SaveBooks(fn, books)
	if err != nil {
		fmt.Println("Error saving books:", err)
		return
	}
	// LOad books from file
	loadedBooks, err := LoadBooks(fn)
	if err != nil {
		fmt.Println("Error loading books:", err)
		return
	}
	fmt.Println("Loaded books:", loadedBooks)
}
