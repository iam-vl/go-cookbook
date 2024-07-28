package main

import "fmt"

func main0204Csv() {
	fn := "books.txt"
	books := []Book{
		{"The Go Programming Language", "Alan A. Donovan", 380},
		{"Go in Action", "Bill Kennedy", 300},
	}
	err := ExportBooksToCsv(fn, books)
	if err != nil {
		fmt.Printf("Failed to export books to CSV: %s\n", err)
	}
	importedBooks, err := ImportBooksFromCsv(fn)
	if err != nil {
		fmt.Printf("Failed to import books from CSV: %s\n", err)
	} else {
		fmt.Printf("Imported books: %+v\n", importedBooks)
	}
}

func main0203Regex() {
	fn := "book_listings.txt"
	books, err := ParseBooksFromFileRegex(fn)
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
