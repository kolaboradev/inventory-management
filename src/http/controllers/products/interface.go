package productController

import "github.com/gofiber/fiber/v2"

type ProductControllerInterface interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	DeleteById(c *fiber.Ctx) error
}
