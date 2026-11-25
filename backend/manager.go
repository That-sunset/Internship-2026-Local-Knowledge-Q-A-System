package main

import (
	"context"

	"internship-2026/db"
	"internship-2026/model"

	"gorm.io/gorm"
)

// --- Record (Session) Managers ---

func getRecordsManager(ctx context.Context) ([]model.Record, error) {
	return gorm.G[model.Record](db.DB).Find(ctx)
}

func createRecordManager(ctx context.Context, record *model.Record) error {
	return gorm.G[model.Record](db.DB).Create(ctx, record)
}

func deleteRecordManager(ctx context.Context, id string) (int, error) {
	return gorm.G[model.Record](db.DB).Where("id = ?", id).Delete(ctx)
}

// --- Chat (Message) Managers ---

func getChatsBySessionManager(ctx context.Context, sessionID string) ([]model.Chat, error) {
	return gorm.G[model.Chat](db.DB).Where("session_id = ?", sessionID).Order("createtime ASC").Find(ctx)
}

func createChatManager(ctx context.Context, chat *model.Chat) error {
	return gorm.G[model.Chat](db.DB).Create(ctx, chat)
}
