package customerController

import "github.com/gofiber/fiber/v2"

type CustomerControllerInterface interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}
