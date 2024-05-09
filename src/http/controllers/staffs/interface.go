package staffController

import "github.com/gofiber/fiber/v2"

type StaffControllerInterface interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
