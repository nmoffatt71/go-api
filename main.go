package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	fmt.Println("Welcome!")
	server := gin.Default()
	server.GET("/events", getEvents)
	server.Run(":8080") // localhost

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "All good!"})

}
