package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	router.Use(cors.Default())

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/tickers", GetAllTickers)
	}

	// WebSocket route
	router.GET("/ws/trades/:ticker", func(c *gin.Context) {
		if c.Request.Header.Get("Upgrade") != "websocket" {
			c.JSON(400, gin.H{"error": "WebSocket upgrade required"})
			return
		}

		// Specific WebSocket logic here
		ListenTicker(c)
	})
}
