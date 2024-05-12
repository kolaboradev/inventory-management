package httpServer

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kolaboradev/inventory/src/helper"
	customerController "github.com/kolaboradev/inventory/src/http/controllers/customers"
	productController "github.com/kolaboradev/inventory/src/http/controllers/products"
	staffController "github.com/kolaboradev/inventory/src/http/controllers/staffs"
	"github.com/kolaboradev/inventory/src/http/middlewares"
	v1Routes "github.com/kolaboradev/inventory/src/http/routes/v1"
	customerV1Routes "github.com/kolaboradev/inventory/src/http/routes/v1/customer"
	productV1Routes "github.com/kolaboradev/inventory/src/http/routes/v1/product"
	staffV1Routes "github.com/kolaboradev/inventory/src/http/routes/v1/staff"
	customerRepository "github.com/kolaboradev/inventory/src/repositories/customer"
	orderRepository "github.com/kolaboradev/inventory/src/repositories/order"
	productRepository "github.com/kolaboradev/inventory/src/repositories/product"
	staffrepository "github.com/kolaboradev/inventory/src/repositories/staff"
	customerService "github.com/kolaboradev/inventory/src/services/customer"
	orderService "github.com/kolaboradev/inventory/src/services/order"
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
	validator.RegisterValidation("url_valid", helper.IsValidUrl)

	// ? Staff
	staffRepo := staffrepository.NewStaffRepository()
	staffService := staffService.NewStaffService(staffRepo, server.DB, validator)
	staffController := staffController.NewStaffController(staffService)

	// ? Order
	orderRepo := orderRepository.NewOrderRepo()
	customerRepo := customerRepository.NewCustomerRepo()

	// ? Product
	productRepo := productRepository.NewProductRepo()
	orderService := orderService.NewOrderService(server.DB, validator, orderRepo, productRepo, customerRepo)
	productService := productService.NewProductService(server.DB, validator, productRepo)
	productController := productController.NewProductController(productService, orderService)

	// ? Customer
	customerService := customerService.NewCustomerService(server.DB, validator, customerRepo)
	customerController := customerController.NewCustomerController(customerService)

	app := fiber.New(fiber.Config{
		ServerHeader: "Kolaboradev",
		ErrorHandler: middlewares.ErrorHandle,
	})

	app.Use(recover.New())
	v1 := v1Routes.SetRoutesV1(app)
	staffV1Routes.SetRoutesStaff(v1, staffController)
	productV1Routes.SetRoutesProduct(v1, productController)
	customerV1Routes.SetRoutesProduct(v1, customerController)

	app.Listen(":8080")
}
