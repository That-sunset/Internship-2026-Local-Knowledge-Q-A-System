package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"internship-2026/model"
)

// RecordData holds only the fields exposed for a session.
type RecordData struct {
	ID         uint       `json:"id"`
	UpdateTime *time.Time `json:"updatetime"`
}

// ChatData holds the fields exposed for an individual chat message.
type ChatData struct {
	ID         uint      `json:"id"`
	SessionID  uint      `json:"session_id"`
	Content    string    `json:"content"`
	Source     string    `json:"source"`
	CreateTime time.Time `json:"createtime"`
}

type ollamaGenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaGenerateResponse struct {
	Response string `json:"response"`
}

func toRecordData(r model.Record) RecordData {
	return RecordData{
		ID:         r.ID,
		UpdateTime: r.UpdateTime,
	}
}

func toChatData(c model.Chat) ChatData {
	return ChatData{
		ID:         c.ID,
		SessionID:  c.SessionID,
		Content:    c.Content,
		Source:     c.Source,
		CreateTime: c.CreateTime,
	}
}

// --- Ollama AI Helper ---

func askModel(ctx context.Context, prompt string) (string, error) {
	reqPayload := ollamaGenerateRequest{
		Model:  "qwen2.5:1.5b",
		Prompt: prompt,
		Stream: false,
	}

	jsonBytes, err := json.Marshal(reqPayload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://127.0.0.1:11434/api/generate", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama returned status %d", resp.StatusCode)
	}

	var result ollamaGenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Response, nil
}

// --- Record (Session) Services ---

func getChatService(ctx context.Context, id string) (RecordData, error) {
	record, err := getChatManager(ctx, id)
	if err != nil {
		return RecordData{}, err
	}
	return toRecordData(record), nil
}

func getRecordsService(ctx context.Context) ([]RecordData, error) {
	records, err := getRecordsManager(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]RecordData, 0, len(records))
	for _, r := range records {
		result = append(result, toRecordData(r))
	}
	return result, nil
}

func createRecordService(ctx context.Context) (RecordData, error) {
	record := model.Record{}
	if err := createRecordManager(ctx, &record); err != nil {
		return RecordData{}, err
	}
	return toRecordData(record), nil
}

func deleteRecordService(ctx context.Context, id string) (int, error) {
	return deleteRecordManager(ctx, id)
}

// --- Chat (Message) Services ---

func getChatsBySessionService(ctx context.Context, sessionID string) ([]ChatData, error) {
	chats, err := getChatsBySessionManager(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	result := make([]ChatData, 0, len(chats))
	for _, c := range chats {
		result = append(result, toChatData(c))
	}
	return result, nil
}

func createChatService(ctx context.Context, sessionID uint, content string, source string) (ChatData, error) {
	chat := model.Chat{
		SessionID: sessionID,
		Content:   content,
		Source:    source,
	}
	if err := createChatManager(ctx, &chat); err != nil {
		return ChatData{}, err
	}
	return toChatData(chat), nil
}

// askAndSaveChatService saves the user prompt, queries Ollama, saves the model reply, and returns the model reply.
func askAndSaveChatService(ctx context.Context, sessionID uint, prompt string) (ChatData, error) {
	// 1. Insert "user message" into SQLite
	if _, err := createChatService(ctx, sessionID, prompt, "user"); err != nil {
		return ChatData{}, err
	}

	// 2. Call askModel(prompt) on Ollama
	answer, err := askModel(ctx, prompt)
	if err != nil {
		return ChatData{}, err
	}

	// 3. Insert "model reply" into SQLite
	modelChat, err := createChatService(ctx, sessionID, answer, "model")
	if err != nil {
		return ChatData{}, err
	}

	return modelChat, nil
}