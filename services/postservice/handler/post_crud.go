package handler

import (
	"net/http"
	db "thesayedirfan/socialapi/utils/db_utils"
	"thesayedirfan/socialapi/utils/http_utils"
	"thesayedirfan/socialapi/utils/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {

	var post models.Post

	if bodyParserErr := c.BodyParser(&post); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	// if exists, _ := userutils.UserExists(db.DB, post.UserID); !exists {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserDoesNotExist})
	// }

	if err := db.DB.Create(&post).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user view", "data": post})
}

func ViewPost(c *fiber.Ctx) error {

	var post models.Post

	if bodyParserErr := c.BodyParser(&post); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	if err := db.DB.Where("id=?", post.ID).First(&post).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	//get comments also
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user view", "data": post})
}

func DeletePost(c *fiber.Ctx) error {

	var post models.Post
	var postID map[string]interface{}

	if bodyParserErr := c.BodyParser(&postID); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	if bodyParserErr := c.BodyParser(&postID); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	ID := postID["id"].(string)

	if ID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserIDError})
	}

	if err := db.DB.Delete(&post, ID).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "post deleted"})
}

func GetUserPost(c *fiber.Ctx) error {

	var userID map[string]interface{}
	var posts []models.Post

	if bodyParserErr := c.BodyParser(&userID); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	ID := userID["id"].(string)

	if ID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserIDError})
	}

	if err := db.DB.Where("user_id = ?", ID).Find(&posts).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user following ", "data": posts})
}
