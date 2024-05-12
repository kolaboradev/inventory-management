package productV1Routes

import (
	"github.com/gofiber/fiber/v2"
	productController "github.com/kolaboradev/inventory/src/http/controllers/products"
	"github.com/kolaboradev/inventory/src/http/middlewares"
)

func SetRoutesProduct(router fiber.Router, pc productController.ProductControllerInterface) {
	productGroup := router.Group("/product")

	productGroup.Get("/customer", pc.FindAllForCustomer)
	productGroup.Post("", middlewares.AuthMiddleware, pc.Create)
	productGroup.Put("/:id", middlewares.AuthMiddleware, pc.Update)
	productGroup.Delete("/:id", middlewares.AuthMiddleware, pc.DeleteById)
	productGroup.Get("", middlewares.AuthMiddleware, pc.FindAll)
	productGroup.Post("/checkout", middlewares.AuthMiddleware, pc.CreateCheckout)
	productGroup.Get("/checkout/history", middlewares.AuthMiddleware, pc.HistoryCheckout)
}
