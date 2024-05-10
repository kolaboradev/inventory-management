package productService

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kolaboradev/inventory/src/exception"
	"github.com/kolaboradev/inventory/src/helper"
	productEntity "github.com/kolaboradev/inventory/src/models/entities/product"
	productRequest "github.com/kolaboradev/inventory/src/models/web/request/product"
	productResponse "github.com/kolaboradev/inventory/src/models/web/response/product"
	productRepository "github.com/kolaboradev/inventory/src/repositories/product"
)

type ProductService struct {
	DB          *sql.DB
	validator   *validator.Validate
	productRepo productRepository.ProductRepositoryInterface
}

func NewProductService(db *sql.DB, validator *validator.Validate, pr productRepository.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{
		DB:          db,
		validator:   validator,
		productRepo: pr,
	}
}

func (service *ProductService) Create(ctx context.Context, request productRequest.ProductCreate) productResponse.ProductCreateResponse {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	id := helper.UUIDStr()
	timeNow := helper.TimeISO8601()

	product := productEntity.Product{
		Id:          id,
		Name:        request.Name,
		Sku:         request.Sku,
		Category:    request.Category,
		ImageUrl:    request.ImageUrl,
		Notes:       request.Notes,
		Price:       request.Price,
		Stock:       request.Stock,
		Location:    request.Location,
		IsAvailable: request.IsAvailable,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}

	service.productRepo.Save(ctx, tx, &product)

	productResponse := productResponse.ProductCreateResponse{
		Id:        product.Id,
		CreatedAt: product.CreatedAt,
	}

	return productResponse
}

func (service *ProductService) Update(ctx context.Context, request productRequest.ProductUpdated) productResponse.ProductUpdateResponse {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	findProduct := service.productRepo.FindById(ctx, tx, request.Id)

	if !findProduct {
		panic(exception.NewNotFoundError("product not found"))
	}

	timeNow := helper.TimeISO8601()

	product := productEntity.Product{
		Id:          request.Id,
		Name:        request.Name,
		Sku:         request.Sku,
		Category:    request.Category,
		ImageUrl:    request.ImageUrl,
		Notes:       request.Notes,
		Price:       request.Price,
		Stock:       request.Stock,
		Location:    request.Location,
		IsAvailable: request.IsAvailable,
		UpdatedAt:   timeNow,
	}

	service.productRepo.Update(ctx, tx, &product)

	return productResponse.ProductUpdateResponse{
		Id:          product.Id,
		Name:        product.Name,
		Sku:         product.Sku,
		Category:    product.Category,
		ImageUrl:    product.ImageUrl,
		Notes:       product.Notes,
		Price:       product.Price,
		Stock:       product.Stock,
		Location:    product.Location,
		IsAvailable: product.IsAvailable,
		UpdateAt:    product.UpdatedAt,
	}
}

func (service *ProductService) DeleteById(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	fmt.Println("service")

	findProduct := service.productRepo.FindById(ctx, tx, id)

	if !findProduct {
		panic(exception.NewNotFoundError("product not found"))
	}

	service.productRepo.DeleteById(ctx, tx, id)
}
