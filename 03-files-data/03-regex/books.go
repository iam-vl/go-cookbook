package main

type Book struct {
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
	Pages  int    `json:"pages" xml:"pages"`
}
