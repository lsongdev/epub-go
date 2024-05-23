package main

import (
	"log"

	"github.com/song940/epub-go/epub"
)

func main() {
	book, err := epub.Open("/Volumes/data/Documents/books/books/11654.epub")
	if err != nil {
		panic(err)
	}
	defer book.Close()
	log.Println(book.Files())
	log.Println(book.Opf.Metadata.Title[0])
	log.Println(book.Opf.Metadata.Description)
}
