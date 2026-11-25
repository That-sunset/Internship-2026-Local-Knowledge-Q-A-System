package main

import (
	"net/http"

	"internship-2026/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	InitVectorDB()

	r := gin.Default()

	// Enable CORS for Vue
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.GET("/api/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/api/records", getHandler)
	r.GET("/api/chat/:id", getChatHandler)
	r.POST("/api/chat/:id", postChatHandler)
	r.POST("/api/records", postHandler)
	r.DELETE("/api/records/:id", deleteHandler)
	r.POST("/api/vector/upload", uploadVectorHandler)

	r.Run()
}
