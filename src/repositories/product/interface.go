package productRepository

import (
	"context"
	"database/sql"

	productEntity "github.com/kolaboradev/inventory/src/models/entities/product"
)

type ProductRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, product *productEntity.Product) *productEntity.Product
	Update(ctx context.Context, tx *sql.Tx, product *productEntity.Product) *productEntity.Product
	FindById(ctx context.Context, tx *sql.Tx, id string) bool
	DeleteById(ctx context.Context, tx *sql.Tx, id string)
}
