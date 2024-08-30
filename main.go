package main

import (
	"example.com/event-registration-app/db"
	"example.com/event-registration-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
