package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kamalesh-seervi/consumer/api"
	"github.com/kamalesh-seervi/consumer/core"
)

func main() {
	core.Load()
	router := gin.Default()

	// CORS middleware
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
