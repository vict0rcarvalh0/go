package main

import (
	"event-booking-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // initialize the default Gin router(with the Logger and Recovery middleware attached)

	server.GET("/events", getEvents)

	server.Run(":8080") // start the server
}

func getEvents(context *gin.Context) { // context is the object that contains all the information about the current request
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
