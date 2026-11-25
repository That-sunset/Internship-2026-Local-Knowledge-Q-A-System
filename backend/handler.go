package main

import (
	"io"
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
		Prompt  string `json:"prompt"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prompt := body.Prompt
	if prompt == "" {
		prompt = body.Content
	}
	if prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prompt is required"})
		return
	}

	modelChat, err := askAndSaveChatService(c.Request.Context(), uint(sessionID), prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"answer":  modelChat.Content,
		"data":    modelChat,
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

func uploadVectorHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid multipart form"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no files uploaded"})
		return
	}

	totalChunks := 0
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rawBytes, readErr := io.ReadAll(file)
		file.Close()
		if readErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": readErr.Error()})
			return
		}

		chunksAdded, err := ingestFileToVectorDB(c.Request.Context(), fileHeader.Filename, rawBytes)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		totalChunks += chunksAdded
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "uploaded",
		"files":   len(files),
		"chunks":  totalChunks,
	})
}
