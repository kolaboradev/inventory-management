package staffController

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/inventory/src/exception"
	staffRequest "github.com/kolaboradev/inventory/src/models/web/request/staff"
	webResponse "github.com/kolaboradev/inventory/src/models/web/response"
	staffService "github.com/kolaboradev/inventory/src/services/staff"
)

type StaffController struct {
	ServiceStaff staffService.StaffServiceInterface
}

func NewStaffController(serviceStaff staffService.StaffServiceInterface) StaffControllerInterface {
	return &StaffController{
		ServiceStaff: serviceStaff,
	}
}

func (controller *StaffController) Register(c *fiber.Ctx) error {
	staffRequest := staffRequest.StaffCreate{}

	if err := c.BodyParser(&staffRequest); err != nil {
		panic(exception.NewBadRequestError("Data invalid"))
	}

	staffResponse := controller.ServiceStaff.Register(context.Background(), staffRequest)

	c.Set("X-Author", "Kolaboradev")
	c.Status(201)

	return c.JSON(webResponse.WebResponse{
		Message: "User registered successfully",
		Data:    staffResponse,
	})
}
func (controller *StaffController) Login(c *fiber.Ctx) error {
	staffRequest := staffRequest.StaffLogin{}
	if err := c.BodyParser(&staffRequest); err != nil {
		panic(exception.NewBadRequestError("Data invalid"))
	}
	staffResponse := controller.ServiceStaff.Login(context.Background(), staffRequest)
	c.Set("X-Author", "Kolaboradev")
	c.Status(200)
	return c.JSON(webResponse.WebResponse{
		Message: "login succes",
		Data:    staffResponse,
	})
}
