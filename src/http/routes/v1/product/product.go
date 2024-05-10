package productV1Routes

import (
	"github.com/gofiber/fiber/v2"
	productController "github.com/kolaboradev/inventory/src/http/controllers/products"
	"github.com/kolaboradev/inventory/src/http/middlewares"
)

func SetRoutesProduct(router fiber.Router, pc productController.ProductControllerInterface) {
	staffGrup := router.Group("/product")

	staffGrup.Post("", middlewares.AuthMiddleware, pc.Create)
	staffGrup.Put("/:id", middlewares.AuthMiddleware, pc.Update)
	staffGrup.Delete("/:id", middlewares.AuthMiddleware, pc.DeleteById)
}
