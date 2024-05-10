
// package customerController

// import (
// 	"context"

// 	"github.com/gofiber/fiber/v2"
// 	customerRequest "github.com/kolaboradev/inventory/src/models/web/request/customer"
// 	webResponse "github.com/kolaboradev/inventory/src/models/web/response"
// 	customerService "github.com/kolaboradev/inventory/src/services/customer"
// )

// type CustomerController struct {
// 	ServiceCustomer customerService.CustomerServiceInterface
// }

// func NewCustomerController(serviceCustomer customerService.CustomerServiceInterface) CustomerControllerInterface {
// 	return &CustomerController{
// 		ServiceCustomer: serviceCustomer,
// 	}
// }

// func (controller *CustomerController) Register(c *fiber.Ctx) error {
// 	customerRequest := customerRequest.CustomerCreate{}

// 	if err := c.BodyParser(&customerRequest); err != nil {
// 		return err
// 	}

// 	customerResponse := controller.ServiceCustomer.Register(context.Background(), customerRequest)

// 	c.Set("X-Author", "Kolaboradev")
// 	c.Status(201)

// 	return c.JSON(webResponse.WebResponse{
// 		Message: "User registered successfully",
// 		Data:    customerResponse,
// 	})
// }
// func (controller *CustomerController) Get(c *fiber.Ctx) error {
// 	panic("bentar bang ")
// func (controller *CustomerController) Checkout(c *fiber.Ctx) error {
// 	panic("bentar bang")
// }
// func (controller *CustomerController) GetHistory(c *fiber.Ctx) error {
// 	panic("bentar bang")
// }