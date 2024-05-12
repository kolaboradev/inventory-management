package productController

import "github.com/gofiber/fiber/v2"

type ProductControllerInterface interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	DeleteById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindAllForCustomer(c *fiber.Ctx) error
	CreateCheckout(c *fiber.Ctx) error
	HistoryCheckout(c *fiber.Ctx) error
}
