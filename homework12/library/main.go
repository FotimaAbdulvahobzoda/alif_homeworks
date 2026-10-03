package main

import (
	"fmt"

	book "library/books"
	"library/catalog"
)

func main() {
	books := []book.Book{
		{
			Title:  "1984",
			Author: "George Orwell",
			Year:   1949,
		},
		{
			Title:  "Animal Farm",
			Author: "George Orwell",
			Year:   1945,
		},
		{
			Title:  "The Hobbit",
			Author: "J.R.R. Tolkien",
			Year:   1937,
		},
		{
			Title:  "The Lord of the Rings",
			Author: "J.R.R. Tolkien",
			Year:   1954,
		},
		{
			Title:  "Dune",
			Author: "Frank Herbert",
			Year:   1965,
		},
	}

	found := catalog.FindByAuthor(books, "George Orwell")

	for _, b := range found {
		fmt.Printf("%s (%d) — %s\n", b.Title, b.Year, b.Author)
	}
}
