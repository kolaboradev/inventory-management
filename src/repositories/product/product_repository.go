package productRepository

import (
	"context"
	"database/sql"

	"github.com/kolaboradev/inventory/src/helper"
	productEntity "github.com/kolaboradev/inventory/src/models/entities/product"
)

type ProductRepository struct {
}

func NewProductRepo() ProductRepositoryInterface {
	return &ProductRepository{}
}

func (repository *ProductRepository) Save(ctx context.Context, tx *sql.Tx, product *productEntity.Product) *productEntity.Product {
	query := "INSERT INTO products (id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"

	_, err := tx.ExecContext(ctx, query, product.Id, product.Name, product.Sku, product.Category, product.ImageUrl, product.Notes, product.Price, product.Stock, product.Location, product.IsAvailable, product.CreatedAt, product.UpdatedAt)
	helper.ErrorIfPanic(err)

	return product
}

func (repository *ProductRepository) Update(ctx context.Context, tx *sql.Tx, product *productEntity.Product) *productEntity.Product {
	query := "UPDATE products SET name = $1, sku = $2, category = $3, image_url = $4, notes = $5, price = $6, stock = $7, location = $8, is_available = $9, updated_at= $10 WHERE id = $11"

	_, err := tx.ExecContext(ctx, query, product.Name, product.Sku, product.Category, product.ImageUrl, product.Notes, product.Price, product.Stock, product.Location, product.IsAvailable, product.UpdatedAt, product.Id)
	helper.ErrorIfPanic(err)

	return product
}

func (repository *ProductRepository) FindById(ctx context.Context, tx *sql.Tx, id string) bool {
	query := "SELECT id FROM products WHERE id = $1"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}

}

func (repository *ProductRepository) DeleteById(ctx context.Context, tx *sql.Tx, id string) {
	query := "DELETE FROM products WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	helper.ErrorIfPanic(err)
}
