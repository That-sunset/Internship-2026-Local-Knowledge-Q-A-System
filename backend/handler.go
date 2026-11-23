package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getHandler(c *gin.Context) {
	records, err := getRecordsService(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func getChatHandler(c *gin.Context) {
	id := c.Param("id")
	chats, err := getChatsBySessionService(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chats)
}

func postHandler(c *gin.Context) {
	record, err := createRecordService(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    record,
	})
}

func postChatHandler(c *gin.Context) {
	idParam := c.Param("id")
	sessionID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	var body struct {
		Content string `json:"content" binding:"required"`
		Source  string `json:"source" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Source != "user" && body.Source != "model" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "source must be 'user' or 'model'"})
		return
	}

	chat, err := createChatService(c.Request.Context(), uint(sessionID), body.Content, body.Source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    chat,
	})
}

func deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if _, err := deleteRecordService(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
		"id":      id,
	})
}