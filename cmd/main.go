package main

import (
	"library-api/pkg/routers"
	"log"
	"net/http"
)

func main() {
	// Configurar el router
	router := routers.SetupRouter()

	// Iniciar el servidor
	log.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
