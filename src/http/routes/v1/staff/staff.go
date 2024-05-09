package staffV1Routes

import (
	"github.com/gofiber/fiber/v2"
	staffController "github.com/kolaboradev/inventory/src/http/controllers/staffs"
)

func SetRoutesStaff(router fiber.Router, sc staffController.StaffControllerInterface) {
	staffGrup := router.Group("/staff")

	staffGrup.Post("/register", sc.Register)
}

// func serve() {
// 	app := fiber.New()
// 	v1 := SetRoutesV1(app)
// 	sc := staffController.StaffController{}
// 	SetRoutesStaff(v1, &sc)

// 	app.Listen(":3000")

// }

// func SetRoutesV1(app *fiber.App) fiber.Router {
// 	v1 := app.Group("/v1")
// 	return v1
// }

// func SetRoutesStaff(router fiber.Router, staffController staffController.StaffControllerInterface) {

// 	staffGroup := router.Group("/staff")

// 	staffGroup.Post("/register", staffController.Login)
// }
