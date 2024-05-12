package productRepository

import (
	"context"
	"database/sql"

	productEntity "github.com/kolaboradev/inventory/src/models/entities/product"
	productRequest "github.com/kolaboradev/inventory/src/models/web/request/product"
)

type ProductRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, product *productEntity.Product) *productEntity.Product
	Update(ctx context.Context, tx *sql.Tx, product *productEntity.Product) *productEntity.Product
	FindByIdBool(ctx context.Context, tx *sql.Tx, id string) bool
	DeleteById(ctx context.Context, tx *sql.Tx, id string)
	FindAll(ctx context.Context, tx *sql.Tx, filters productRequest.ProductGetFilter) []productEntity.Product
	FindById(ctx context.Context, tx *sql.Tx, id string) productEntity.Product
}
