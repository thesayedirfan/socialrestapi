package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	PostID    uint      `json:"post_id"`
	Comment   string    `json:"comment"`
	UserID    uint      `gorm:"index" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
