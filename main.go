package main

import (
	"log"
	"os"

	"github.com/lsongdev/epub-go/epub"
)

func main() {
	book, err := epub.Open("/Volumes/data/Documents/books/books/25828.epub")
	if err != nil {
		panic(err)
	}
	defer book.Close()
	// log.Println(book.Files())
	log.Println(book.Title())
	log.Println(book.Author())
	// log.Println(book.Opf.Metadata.Description)
	cover, err := book.ReadCover()
	if err == nil {
		os.WriteFile(book.Title()+".png", cover, 0644)
	}
}
