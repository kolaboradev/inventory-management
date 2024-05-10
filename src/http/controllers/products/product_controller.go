package productController

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	productRequest "github.com/kolaboradev/inventory/src/models/web/request/product"
	webResponse "github.com/kolaboradev/inventory/src/models/web/response"
	productService "github.com/kolaboradev/inventory/src/services/product"
)

type ProductController struct {
	productService productService.ProductServiceInterface
}

func NewProductController(ps productService.ProductServiceInterface) ProductControllerInterface {
	return &ProductController{
		productService: ps,
	}
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	productRequest := productRequest.ProductCreate{}
	staff := c.Locals("staffId").(string)
	fmt.Println(staff)
	if err := c.BodyParser(&productRequest); err != nil {
		return err
	}
	productResponse := controller.productService.Create(context.Background(), productRequest)

	c.Set("X-Author", "Kolaboradev")
	c.Status(201)

	return c.JSON(webResponse.WebResponse{
		Message: "Create new product succes",
		Data:    productResponse,
	})
}

func (controller *ProductController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	productRequest := productRequest.ProductUpdated{
		Id: id,
	}
	if err := c.BodyParser(&productRequest); err != nil {
		return err
	}
	productResponse := controller.productService.Update(context.Background(), productRequest)

	c.Set("X-Author", "Kolaboradev")
	return c.JSON(webResponse.WebResponse{
		Message: "Update product succes",
		Data:    productResponse,
	})
}

func (controller *ProductController) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	controller.productService.DeleteById(context.Background(), id)
	c.Set("X-Author", "Kolaboradev")
	return c.JSON(webResponse.WebResponse{
		Message: "Delete product succes",
		Data:    "OK",
	})
}
