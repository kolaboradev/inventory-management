package middlewares

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/inventory/src/exception"
	"github.com/kolaboradev/inventory/src/helper"
	webResponse "github.com/kolaboradev/inventory/src/models/web/response"
)

func ErrorHandle(c *fiber.Ctx, err error) error {
	if value, ok := err.(*exception.ConflictError); ok {
		valErr := value.Error()
		c.Status(409)
		return c.JSON(webResponse.WebResponse{
			Message: "CONFLICT ERROR",
			Data:    valErr,
		})
	}
	if value, ok := err.(*exception.BadRequestError); ok {
		valErr := value.Error()
		c.Status(400)
		return c.JSON(webResponse.WebResponse{
			Message: "BAD REQUEST",
			Data:    valErr,
		})
	}
	if value, ok := err.(*exception.NotFoundError); ok {
		valErr := value.Error()
		c.Status(404)
		return c.JSON(webResponse.WebResponse{
			Message: "NOT FOUND",
			Data:    valErr,
		})
	}
	if value, ok := err.(validator.ValidationErrors); ok {
		valErr := value[0]
		message := helper.CustomMessageValidation(valErr)
		c.Status(400)
		return c.JSON(webResponse.WebResponse{
			Message: "Validation Error",
			Data:    message,
		})
	}

	fmt.Println(err)
	c.Status(500)
	return c.JSON(webResponse.WebResponse{
		Message: "SERVER ERROR",
		Data:    err.Error(),
	})
}
