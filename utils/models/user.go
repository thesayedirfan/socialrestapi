package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	UserName  string    `gorm:"unique;index" json:"username"`
	Email     string    `gorm:"unique;index" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
