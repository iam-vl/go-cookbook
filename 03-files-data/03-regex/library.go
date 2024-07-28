package main

import "encoding/xml"

// For XML, we often use a wrapper type for a collection
type Library struct {
	Books []Book `xml:"book"`
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
