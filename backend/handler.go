package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHandler(c *gin.Context) {
	records := []gin.H{
		{"id": 1, "content": "one"},
		{"id": 2, "content": "two"},
		{"id": 3, "content": "three"},
	}
	c.JSON(http.StatusOK, records)
}

func getChatHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"content": "chat content",
	})
}

func postHandler(c *gin.Context) {
	var body gin.H
	c.BindJSON(&body)
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    body,
	})
}

func deleteHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
		"id":      id,
	})
}