package main

import (
	"event-booking-api/db"
	"event-booking-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()                // initialize the database
	server := gin.Default()    // initialize the default gin router(with the logger and recovery middleware attached)]
	router.SetupRoutes(server) // setup the routes

	server.Run(":8080") // start the server
}
