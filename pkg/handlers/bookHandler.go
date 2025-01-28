package handlers

import (
	"encoding/json"
	"library-api/pkg/db"
	"library-api/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := db.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	db.AddBook(book)
	json.NewEncoder(w).Encode(book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := db.GetBookByID(params["id"])
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	db.UpdateBook(params["id"], book)
	json.NewEncoder(w).Encode(book)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.DeleteBook(params["id"])
	w.WriteHeader(http.StatusNoContent)
}
