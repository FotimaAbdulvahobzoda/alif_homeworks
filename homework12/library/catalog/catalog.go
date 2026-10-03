package catalog

import book "library/books"

func FindByAuthor(books []book.Book, author string) []book.Book {
	newSlice := []book.Book{}
	for _, b := range books {
		if author == b.Author {
			newSlice = append(newSlice, b)
		}
	}
	return newSlice
}
