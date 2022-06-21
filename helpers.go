package main

import "fmt"

func verifyBook(newBook book) (book, error) {
	if newBook.Author == "" || newBook.ID == "" || newBook.Title == "" {
		return newBook, fmt.Errorf("malformed Json")
	}
	return newBook, nil
}
