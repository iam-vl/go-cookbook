package main

import "fmt"

func main0207ArrangeFolders() {
	libPath := "../00-library"
	books := Get2BookFiles()
	err := ArrangeBooksByAuthor(libPath, books)
	if err != nil {
		fmt.Printf("Failed arranging books by author: %s\n", err)
	}
}

func main0206BasicReport() {
	books := Get2Books()
	GenerateLibSummary(books)
	err := ExportLibDataForAnalysis("lib_analysis.csv", books)
	if err != nil {
		fmt.Printf("Failed to export lib data for analysis: %s\n", err)
	}

}

func main0205CopyBinaryCover() {
	originalCoverPath := "hm.jpg"
	newCoverPath := "hm2.jpg"
	cover, err := ReadCoverImage(originalCoverPath)
	if err != nil {
		fmt.Println("Failed to read cover image: %s\n", err)
		return
	}
	err = WriteCoverImage(newCoverPath, cover)
	if err != nil {
		fmt.Println("Failed to write cover image: %s\n", err)
	}
}

func main0204Csv() {
	fn := "books.txt"
	books := Get2Books()
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
	books := Get2Books()
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
	books := Get2Books()
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
