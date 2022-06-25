package handler

import (
	"net/http"
	db "thesayedirfan/socialapi/utils/db_utils"
	"thesayedirfan/socialapi/utils/http_utils"
	"thesayedirfan/socialapi/utils/models"

	"github.com/gofiber/fiber/v2"
)

func CommentToPost(c *fiber.Ctx) error {
	var comment models.Comment

	if bodyParserErr := c.BodyParser(&comment); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}
	exists, _, userID := db.PostExists(db.DB, comment.PostID)

	if !exists {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.PostDoesNotExist})
	}

	comment.UserID = userID

	if err := db.DB.Create(&comment).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "comment created", "data": comment})
}

func DeleteComment(c *fiber.Ctx) error {

	var comment models.Comment
	var commentID map[string]interface{}

	if bodyParserErr := c.BodyParser(&commentID); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	ID := commentID["id"].(string)

	if ID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.CommentIDError})
	}
	if err := db.DB.Delete(&comment, ID).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "comment deleted"})
}

func GetPostComments(c *fiber.Ctx) error {

	var comment []models.Comment
	var request map[string]interface{}

	if bodyParserErr := c.BodyParser(&request); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	postID := uint(request["id"].(float64))

	exists, _, _ := db.PostExists(db.DB, postID)

	if !exists {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.PostDoesNotExist})
	}

	if err := db.DB.Where("post_id = ?", postID).Find(&comment).Error; err != nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "comments", "data": comment})

}
