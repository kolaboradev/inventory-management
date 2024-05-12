package customerController

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/inventory/src/exception"
	customerRequest "github.com/kolaboradev/inventory/src/models/web/request/customer"
	webResponse "github.com/kolaboradev/inventory/src/models/web/response"
	customerService "github.com/kolaboradev/inventory/src/services/customer"
)

type CustomerController struct {
	customerService customerService.CustomerServiceInterface
}

func NewCustomerController(cs customerService.CustomerServiceInterface) CustomerControllerInterface {
	return &CustomerController{
		customerService: cs,
	}
}

func (controller *CustomerController) Create(c *fiber.Ctx) error {
	customerRequest := customerRequest.CustomerCreateRequest{}
	if err := c.BodyParser(&customerRequest); err != nil {
		panic(exception.NewBadRequestError("Data invalid"))
	}

	customerResponse := controller.customerService.Create(context.Background(), customerRequest)

	c.Set("X-Author", "Kolaboradev")
	c.Status(201)

	return c.JSON(webResponse.WebResponse{
		Message: "Register Customer product succes",
		Data:    customerResponse,
	})
}

func (controller *CustomerController) FindAll(c *fiber.Ctx) error {
	customerRequest := customerRequest.CustomerFilter{
		PhoneNumber: c.Query("phoneNumber"),
		Name:        c.Query("name"),
	}

	customerResponses := controller.customerService.FindAll(context.Background(), customerRequest)

	if len(customerResponses) == 0 {
		c.Set("X-Author", "Kolaboradev")
		return c.JSON(webResponse.WebResponse{
			Message: "Get all customer succes",
			Data:    []interface{}{},
		})
	}

	c.Set("X-Author", "Kolaboradev")
	return c.JSON(webResponse.WebResponse{
		Message: "Get all customer succes",
		Data:    customerResponses,
	})
}
