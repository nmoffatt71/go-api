package main

import (
	"github.com/gin-gonic/gin"

	"rest-api.com/m/v2/db"
	"rest-api.com/m/v2/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost

}
