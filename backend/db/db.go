package db

import (
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Record struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Content    string         `gorm:"not null" json:"content"`
	CreateTime time.Time      `gorm:"column:createtime;autoCreateTime" json:"createtime"`
	UpdateTime *time.Time     `gorm:"column:updatetime;default:null" json:"updatetime"`
	DeleteTime gorm.DeletedAt `gorm:"column:deletetime;index" json:"deletetime"`
}

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("db/chat.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Record{})
	DB = db
}