package httpServer

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	staffController "github.com/kolaboradev/inventory/src/http/controllers/staffs"
	v1Routes "github.com/kolaboradev/inventory/src/http/routes/v1"
	staffV1Routes "github.com/kolaboradev/inventory/src/http/routes/v1/staff"
	staffrepository "github.com/kolaboradev/inventory/src/repositories/staff"
	staffService "github.com/kolaboradev/inventory/src/services/staff"
)

type HttpServer struct {
	DB *sql.DB
}

func NewServer(db *sql.DB) HttpServerInterface {
	return &HttpServer{DB: db}
}

func (server *HttpServer) Listen() {
	validator := validator.New()

	staffRepo := staffrepository.NewStaffRepository()
	staffService := staffService.NewStaffService(staffRepo, server.DB, *validator)
	staffController := staffController.NewStaffController(staffService)

	app := fiber.New(fiber.Config{
		ServerHeader: "Kolaboradev",
	})
	v1 := v1Routes.SetRoutesV1(app)
	staffV1Routes.SetRoutesStaff(v1, staffController)

	app.Listen(":3000")
}
