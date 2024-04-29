package main

import (
	"account-manager-api/internal/routes"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load()

    r := routes.SetupRoutes()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    address := fmt.Sprintf(":%s", port)

    log.Printf("Servidor rodando na porta %v", port)

    log.Fatal(http.ListenAndServe(address, r))
}
