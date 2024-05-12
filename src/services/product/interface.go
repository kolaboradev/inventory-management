package productService

import (
	"context"

	productRequest "github.com/kolaboradev/inventory/src/models/web/request/product"
	productResponse "github.com/kolaboradev/inventory/src/models/web/response/product"
)

type ProductServiceInterface interface {
	Create(ctx context.Context, request productRequest.ProductCreate) productResponse.ProductCreateResponse
	Update(ctx context.Context, request productRequest.ProductUpdated) productResponse.ProductUpdateResponse
	DeleteById(ctx context.Context, id string)
	FindAll(ctx context.Context, filters productRequest.ProductGetFilter) []productResponse.ProductGetResponse
	FindAllForCustomer(ctx context.Context, filters productRequest.ProductGetFilter) []productResponse.ProductGetResponse
}
