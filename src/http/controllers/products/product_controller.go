package productController

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/inventory/src/exception"
	orderRequest "github.com/kolaboradev/inventory/src/models/web/request/order"
	productRequest "github.com/kolaboradev/inventory/src/models/web/request/product"
	webResponse "github.com/kolaboradev/inventory/src/models/web/response"
	orderService "github.com/kolaboradev/inventory/src/services/order"
	productService "github.com/kolaboradev/inventory/src/services/product"
)

type ProductController struct {
	productService productService.ProductServiceInterface
	orderService   orderService.OrderServiceInterface
}

func NewProductController(ps productService.ProductServiceInterface, os orderService.OrderServiceInterface) ProductControllerInterface {
	return &ProductController{
		productService: ps,
		orderService:   os,
	}
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	productRequest := productRequest.ProductCreate{}
	if err := c.BodyParser(&productRequest); err != nil {
		panic(exception.NewBadRequestError("Data invalid"))
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
		panic(exception.NewBadRequestError("Data invalid"))
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
	controller.productService.DeleteById(context.Background(), id)
	c.Set("X-Author", "Kolaboradev")
	return c.JSON(webResponse.WebResponse{
		Message: "Delete product succes",
		Data:    "OK",
	})
}

func (controller *ProductController) FindAll(c *fiber.Ctx) error {
	productRequest := productRequest.ProductGetFilter{
		Id:          c.Query("id", ""),
		Limit:       c.QueryInt("limit", 0),
		Offset:      c.QueryInt("offset", 0),
		Name:        c.Query("name", ""),
		IsAvailable: c.Query("isAvailable"),
		Category:    c.Query("category", ""),
		Sku:         c.Query("sku", ""),
		Price:       c.Query("price", ""),
		InStock:     c.Query("inStock", ""),
		CreatedAt:   c.Query("createdAt"),
	}

	productResponses := controller.productService.FindAll(context.Background(), productRequest)

	if len(productResponses) == 0 {
		c.Set("X-Author", "Kolaboradev")
		return c.JSON(webResponse.WebResponse{
			Message: "Get all product succes",
			Data:    []interface{}{},
		})
	}

	c.Set("X-Author", "Kolaboradev")
	return c.JSON(webResponse.WebResponse{
		Message: "Get all product succes",
		Data:    productResponses,
	})
}

func (controller *ProductController) FindAllForCustomer(c *fiber.Ctx) error {
	productRequest := productRequest.ProductGetFilter{
		Limit:       c.QueryInt("limit", 0),
		Offset:      c.QueryInt("offset", 0),
		Name:        c.Query("name", ""),
		IsAvailable: "true",
		Category:    c.Query("category", ""),
		Sku:         c.Query("sku", ""),
		Price:       c.Query("price", ""),
		InStock:     c.Query("inStock", ""),
	}

	productResponses := controller.productService.FindAll(context.Background(), productRequest)

	if len(productResponses) == 0 {
		c.Set("X-Author", "Kolaboradev")
		return c.JSON(webResponse.WebResponse{
			Message: "Get all product succes",
			Data:    []interface{}{},
		})
	}

	c.Set("X-Author", "Kolaboradev")
	return c.JSON(webResponse.WebResponse{
		Message: "Get all product succes",
		Data:    productResponses,
	})
}

func (controller *ProductController) CreateCheckout(c *fiber.Ctx) error {
	orderRequest := orderRequest.OrderCreateRequest{}
	if err := c.BodyParser(&orderRequest); err != nil {
		panic(exception.NewBadRequestError("Data invalid"))
	}

	orderResponse := controller.orderService.Create(context.Background(), orderRequest)

	return c.JSON(webResponse.WebResponse{
		Message: "succes checkout product",
		Data:    orderResponse,
	})
}

func (controller *ProductController) HistoryCheckout(c *fiber.Ctx) error {
	filterRequest := orderRequest.OrderGet{
		CustomerId: c.Query("customerId", ""),
		CreatedAt:  c.Query("createdAt", ""),
		Limit:      c.QueryInt("limit", 0),
		Offset:     c.QueryInt("offset", 0),
	}

	orderResponses := controller.orderService.FindAll(context.Background(), filterRequest)
	if len(orderResponses) == 0 {
		c.Set("X-Author", "Kolaboradev")
		return c.JSON(webResponse.WebResponse{
			Message: "Get all history transaction succes",
			Data:    []interface{}{},
		})
	}
	c.Set("X-Author", "Kolaboradev")
	return c.JSON(webResponse.WebResponse{
		Message: "Get all history transaction succes",
		Data:    orderResponses,
	})
}
