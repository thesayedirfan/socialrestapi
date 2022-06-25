package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Body      string    `json:"body"`
	UserID    uint      `gorm:"index" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
