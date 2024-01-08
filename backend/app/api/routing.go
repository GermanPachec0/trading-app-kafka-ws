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
	router.GET("/ws/trades/:ticker", func(ctx *gin.Context) {
		if ctx.Request.Header.Get("Upgrade") != "websocket" {
			ctx.JSON(400, gin.H{"error": "WebSocket upgrade required"})
			return
		}
		ListenTicker(ctx)
	})
}
