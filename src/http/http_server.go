package httpServer

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kolaboradev/inventory/src/helper"
	productController "github.com/kolaboradev/inventory/src/http/controllers/products"
	staffController "github.com/kolaboradev/inventory/src/http/controllers/staffs"
	"github.com/kolaboradev/inventory/src/http/middlewares"
	v1Routes "github.com/kolaboradev/inventory/src/http/routes/v1"
	productV1Routes "github.com/kolaboradev/inventory/src/http/routes/v1/product"
	staffV1Routes "github.com/kolaboradev/inventory/src/http/routes/v1/staff"
	productRepository "github.com/kolaboradev/inventory/src/repositories/product"
	staffrepository "github.com/kolaboradev/inventory/src/repositories/staff"
	productService "github.com/kolaboradev/inventory/src/services/product"
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
	validator.RegisterValidation("phone_number", helper.MustPhoneNumber)
	validator.RegisterValidation("category", helper.MatchCategoryProduct)

	// ? Staff
	staffRepo := staffrepository.NewStaffRepository()
	staffService := staffService.NewStaffService(staffRepo, server.DB, validator)
	staffController := staffController.NewStaffController(staffService)

	// ? Product
	productRepo := productRepository.NewProductRepo()
	productService := productService.NewProductService(server.DB, validator, productRepo)
	productController := productController.NewProductController(productService)

	app := fiber.New(fiber.Config{
		ServerHeader: "Kolaboradev",
		ErrorHandler: middlewares.ErrorHandle,
	})

	app.Use(recover.New())
	v1 := v1Routes.SetRoutesV1(app)
	staffV1Routes.SetRoutesStaff(v1, staffController)
	productV1Routes.SetRoutesProduct(v1, productController)

	app.Listen(":3000")
}
