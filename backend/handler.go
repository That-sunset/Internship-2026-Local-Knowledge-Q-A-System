package main

import (
	"net/http"

	"internship-2026/db"

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
	var record db.Record
	if err := db.DB.First(&record, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, record)
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