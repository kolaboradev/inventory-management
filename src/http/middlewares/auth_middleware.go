package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kolaboradev/inventory/src/helper"
	webResponse "github.com/kolaboradev/inventory/src/models/web/response"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(webResponse.WebResponse{
			Message: "UNAUTHORIZED",
			Data:    "Token Invalid",
		})
	}
	token, err := jwt.Parse(tokenStr, helper.CheckTokenJWT)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(webResponse.WebResponse{
			Message: "UNAUTHORIZED",
			Data:    "Token Invalid",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(webResponse.WebResponse{
			Message: "UNAUTHORIZED",
			Data:    "Token Invalid",
		})
	}

	c.Locals("staffId", token.Claims.(jwt.MapClaims)["staffId"])

	return c.Next()
}
