package db

import (
	"library-api/pkg/models"
)

var books []models.Book

func GetAllBooks() []models.Book {
	return books
}

func AddBook(book models.Book) {
	books = append(books, book)
}

func GetBookByID(id string) *models.Book {
	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func UpdateBook(id string, updatedBook models.Book) {
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			return
		}
	}
}

func DeleteBook(id string) {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return
		}
	}
}
