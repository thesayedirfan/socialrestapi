package handler

import (
	"net/http"
	"thesayedirfan/socialapi/utils/db_utils"
	"thesayedirfan/socialapi/utils/http_utils"
	"thesayedirfan/socialapi/utils/models"

	"github.com/gofiber/fiber/v2"
)

func UserView(c *fiber.Ctx) error {

	var user models.User

	if bodyParserErr := c.BodyParser(&user); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	if err := db.DB.Where("id=?", user.ID).First(&user).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": err.Error()})
	}

	user.Password = ""

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user view", "data": user})
}

func UserList(c *fiber.Ctx) error {

	var users []models.User

	db.DB.Find(&users)

	for _, user := range users {
		user.Password = ""
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user list", "data": users})
}
