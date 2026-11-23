package model

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreateTime time.Time      `gorm:"column:createtime;autoCreateTime" json:"createtime"`
	UpdateTime *time.Time     `gorm:"column:updatetime;default:null" json:"updatetime"`
	DeleteTime gorm.DeletedAt `gorm:"column:deletetime;index" json:"deletetime"`
}

type Chat struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionID  uint      `gorm:"column:session_id;index;not null" json:"session_id"`
	Content    string    `gorm:"column:content;type:text;not null" json:"content"`
	Source     string    `gorm:"column:source;type:varchar(20);not null" json:"source"`
	CreateTime time.Time `gorm:"column:createtime;autoCreateTime" json:"createtime"`
}