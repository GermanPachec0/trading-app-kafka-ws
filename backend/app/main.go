package main

import (
	"github.com/GermanPachec0/app/api"
	"github.com/GermanPachec0/app/core"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.Load()
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Content-Length", "Accept-Language", "Accept-Encoding", "Connection", "Access-Control-Allow-Origin"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	// Add routes

	api.AddRoutes(router)

	// Start server
	router.Run(":8000")
}
