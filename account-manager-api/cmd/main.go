package main

import (
	"log"
	"net/http"

	"account-manager-api/internal/routes"
)

func main() {
	r := routes.SetupRoutes()

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
