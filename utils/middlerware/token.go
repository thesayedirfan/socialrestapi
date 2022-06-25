package middlerware

import (
	"fmt"
	"net/http"
	"thesayedirfan/socialapi/utils/http_utils"

	"github.com/gofiber/fiber/v2"
)

func TokenValid(ctx *fiber.Ctx) error {
	_, err := VerifyToken(ctx.GetReqHeaders())
	fmt.Println(err)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": http_utils.HttpError.Error, "message": "not authrorized access"})
	}
	ctx.Next()
	return nil
}
