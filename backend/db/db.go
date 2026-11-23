package db

import (
	"log"

	"internship-2026/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("db/chat.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&model.Record{}, &model.Chat{}); err != nil {
		log.Fatal(err)
	}
	if db.Migrator().HasColumn(&model.Record{}, "content") {
		if err := db.Migrator().DropColumn(&model.Record{}, "content"); err != nil {
			log.Fatal(err)
		}
	}
	DB = db
}