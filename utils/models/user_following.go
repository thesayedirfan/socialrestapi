package models

import "time"

type UserFollows struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"index" json:"user_id"`
	FollowingID uint      `gorm:"index" json:"following_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
