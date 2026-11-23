package main

import (
	"context"
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