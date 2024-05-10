package staffV1Routes

import (
	"github.com/gofiber/fiber/v2"
	staffController "github.com/kolaboradev/inventory/src/http/controllers/staffs"
)

func SetRoutesStaff(router fiber.Router, sc staffController.StaffControllerInterface) {
	staffGrup := router.Group("/staff")

	staffGrup.Post("/register", sc.Register)
	staffGrup.Post("/login", sc.Login)
}
