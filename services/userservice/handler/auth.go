package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	db "thesayedirfan/socialapi/utils/db_utils"
	"thesayedirfan/socialapi/utils/http_utils"
	"thesayedirfan/socialapi/utils/models"
	"thesayedirfan/socialapi/utils/redis"
	utils "thesayedirfan/socialapi/utils/service_utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {

	user := models.User{}

	if bodyParserErr := c.BodyParser(&user); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	hashedPassword, passwordErr := utils.HashPassword(user.Password)

	if passwordErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.PasswordParseError, "error": passwordErr})
	}

	user.Password = string(hashedPassword)

	tx := db.DB.Begin()

	userCreateError := tx.Create(&user).Error

	if userCreateError != nil {
		tx.Rollback()
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.DbError, "error": userCreateError})
	}

	tx.Commit()

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "user created", "data": user})
}

func SignIn(c *fiber.Ctx) error {

	var user models.User
	var userResult models.User
	var token *models.Token

	if bodyParserErr := c.BodyParser(&user); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": "unable to parse body", "error": bodyParserErr})
	}

	if err := db.DB.Where("user_name=?", user.UserName).First(&userResult).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": err.Error()})
	}

	passwordMatch := utils.CheckPassword([]byte(user.Password), []byte(userResult.Password))

	if !passwordMatch {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UnableToLogin})
	}

	token, _ = redis.CreateToken(uint64(userResult.ID))

	if TokenErr := redis.SaveTokenInRedis(uint64(userResult.ID), token); TokenErr != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.UnableToLogin, "error": TokenErr})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "login succesfull", "data": userResult, "tokens": token})
}

func Refresh(c *fiber.Ctx) error {
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
	mapToken := map[string]string{}

	if bodyParserErr := c.BodyParser(&mapToken); bodyParserErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": bodyParserErr})
	}

	refreshToken := mapToken["refresh_token"]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.SomethingWentWrong, "error": err})
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.SomethingWentWrong, "error": err})
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		fmt.Println(refreshUuid)
		if !ok {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.SomethingWentWrong, "error": "refresh token"})
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.SomethingWentWrong, "error": "no user"})

		}

		_, delErr := redis.DeleteUuid(refreshUuid)
		if delErr != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.SomethingWentWrong, "error": delErr})
		}

		ts, createErr := redis.CreateToken(userId)
		if createErr != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.SomethingWentWrong, "error": err})
		}

		saveErr := redis.SaveTokenInRedis(userId, ts)
		if saveErr != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.SomethingWentWrong, "error": err})
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": http_utils.HttpError.Success, "message": "tokens  success fully", "tokens": tokens})
	} else {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": http_utils.ErrorMessages.BodyParseError, "error": err})
	}
}
