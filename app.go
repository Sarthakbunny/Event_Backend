package main

import (
	"events.com/m/db"
	"events.com/m/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}
