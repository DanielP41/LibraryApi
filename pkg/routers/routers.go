package routers

import (
	"library-api/pkg/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Definir las rutas
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/book", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", handlers.DeleteBook).Methods("DELETE")

	return router
}
