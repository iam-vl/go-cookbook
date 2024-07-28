package main

import (
	"io"
	"os"
)

type Book struct {
	Title     string `json:"title" xml:"title"`
	Author    string `json:"author" xml:"author"`
	Pages     int    `json:"pages" xml:"pages"`
	Coverpath string // path
}

func WriteCoverImage(filepath string, data []byte) error {
	return os.WriteFile(filepath, data, 0644)
}

func ReadCoverImage(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	imageData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return imageData, nil

}
