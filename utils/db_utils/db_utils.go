package db

import (
	"thesayedirfan/socialapi/utils/models"

	"gorm.io/gorm"
)

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.UserFollows{})
}

func UserExists(db *gorm.DB, id uint) (bool, uint) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return false, 0
	}
	return true, user.ID
}

func PostExists(db *gorm.DB, id uint) (bool, uint, uint) {
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		return false, 0, 0
	}
	return true, post.ID, post.UserID
}
