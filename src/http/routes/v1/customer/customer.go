package customerV1Routes

import (
	"github.com/gofiber/fiber/v2"
	customerController "github.com/kolaboradev/inventory/src/http/controllers/customers"
	"github.com/kolaboradev/inventory/src/http/middlewares"
)

func SetRoutesProduct(router fiber.Router, cc customerController.CustomerControllerInterface) {
	customerGroup := router.Group("/customer")

	customerGroup.Get("", middlewares.AuthMiddleware, cc.FindAll)
	customerGroup.Post("/register", middlewares.AuthMiddleware, cc.Create)
}
