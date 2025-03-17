package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"net/http"

	"rest-api.com/m/v2/models"
)

func main() {
	fmt.Println("Welcome!")
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad data/missing"})
		return
	}

	event.ID = 1
	event.UserId = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Created!", "event": event})

}
