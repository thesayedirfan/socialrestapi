package middlerware

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"thesayedirfan/socialapi/utils/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func VerifyToken(header map[string]string) (*jwt.Token, error) {
	tokenString := header["Authorization"]
	if tokenString == "" {
		return nil, errors.New("No Authtoken")
	}
	tokenArr := strings.Split(tokenString, " ")
	fmt.Println(tokenArr[1])
	token, err := jwt.Parse(tokenArr[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*models.AccessDetails, error) {
	token, err := VerifyToken(c.GetReqHeaders())
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &models.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}
