package handler

import (
	"errors"
	"net/http"
	db "thesayedirfan/socialapi/utils/db_utils"
	"thesayedirfan/socialapi/utils/http_utils"
	"thesayedirfan/socialapi/utils/models"

	"github.com/gofiber/fiber/v2"
)

func UserFollow(c *fiber.Ctx) error {
	var userFollow models.UserFollows

	if bodyParserErr := c.BodyParser(&userFollow); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	if userFollow.UserID == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserDoesNotExist, "error": errors.New("user id not givem")})
	}
	if userFollow.FollowingID == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserDoesNotExist, "error": errors.New("user id not givem")})
	}

	if ok, _ := db.UserExists(db.DB, userFollow.UserID); !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserDoesNotExist, "error": errors.New("user id not givem")})
	}

	if ok, _ := db.UserExists(db.DB, userFollow.FollowingID); !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserDoesNotExist, "error": errors.New("user id not givem")})
	}

	if err := db.DB.Create(&userFollow).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user list", "data": userFollow})
}

func UserUnFollow(c *fiber.Ctx) error {

	var userFollowID map[string]interface{}
	var userFollow models.UserFollows

	if bodyParserErr := c.BodyParser(&userFollowID); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	ID := userFollowID["id"].(string)

	if ID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserIDError})
	}

	if err := db.DB.Delete(&userFollow, ID).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user following deleted"})
}

func GetUserFollowing(c *fiber.Ctx) error {

	var userFollowID map[string]interface{}
	var userFollow []models.UserFollows

	if bodyParserErr := c.BodyParser(&userFollowID); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	ID := userFollowID["id"].(string)

	if ID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UserIDError})
	}

	if err := db.DB.Where("user_id = ?", ID).Find(&userFollow).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user following ", "data": userFollow})
}
