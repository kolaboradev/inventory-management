package productRepository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/kolaboradev/inventory/src/helper"
	productEntity "github.com/kolaboradev/inventory/src/models/entities/product"
	productRequest "github.com/kolaboradev/inventory/src/models/web/request/product"
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

func (repository *ProductRepository) FindByIdBool(ctx context.Context, tx *sql.Tx, id string) bool {
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

func (repository *ProductRepository) FindAll(ctx context.Context, tx *sql.Tx, filters productRequest.ProductGetFilter) []productEntity.Product {
	query := "SELECT id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at, updated_at FROM products WHERE 1 = 1"

	var args []interface{}
	argIndex := 1

	if filters.Id != "" {
		query += fmt.Sprintf(" AND id = $%d", argIndex)
		args = append(args, filters.Id)
		argIndex++
	}
	if filters.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, filters.Name)
		argIndex++
	}
	if filters.IsAvailable != nil {
		valStr, _ := filters.IsAvailable.(string)
		if strings.ToLower(valStr) == "true" || strings.ToLower(valStr) == "false" {
			booleanValue, _ := strconv.ParseBool(valStr)

			filters.IsAvailable = booleanValue

			if _, ok := filters.IsAvailable.(bool); ok {
				query += fmt.Sprintf(" AND is_available = $%d", argIndex)
				args = append(args, filters.IsAvailable)
				argIndex++
			}
		}

	}
	if filters.Category != "" {
		query += fmt.Sprintf(" AND category = $%d", argIndex)
		args = append(args, filters.Category)
		argIndex++
	}
	if filters.Sku != "" {
		query += fmt.Sprintf(" AND sku = $%d", argIndex)
		args = append(args, filters.Sku)
		argIndex++
	}
	if filters.InStock != nil {
		valStr, _ := filters.InStock.(string)
		if strings.ToLower(valStr) == "true" || strings.ToLower(valStr) == "false" {
			booleanValue, _ := strconv.ParseBool(filters.InStock.(string))
			filters.InStock = booleanValue
			if _, ok := filters.InStock.(bool); ok {
				if filters.InStock == true {
					query += " AND stock > 0"
				} else if filters.InStock == false {
					query += " AND stock = 0"
				}
			}

		}
	}

	// query += " ORDER BY created_at DESC"
	var orderBy []string

	if (filters.Price == "asc" || filters.Price == "desc") || (filters.CreatedAt == "asc" || filters.CreatedAt == "desc") {
		query += " ORDER BY"
		if filters.Price == "asc" {
			orderBy = append(orderBy, " price ASC")
		}
		if filters.Price == "desc" {
			orderBy = append(orderBy, " price DESC")
		}
		if filters.CreatedAt == "asc" {
			orderBy = append(orderBy, " created_at ASC")
		}
		if filters.CreatedAt == "desc" {
			orderBy = append(orderBy, " created_at DESC")
		}

		result := strings.Join(orderBy, ",")
		query += result

	}

	if len(orderBy) == 0 {
		query += " ORDER BY created_at DESC"
	}

	if filters.Offset >= 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, filters.Offset)
		argIndex++
	}
	if filters.Limit >= 0 {
		if filters.Limit == 0 {
			filters.Limit = 5
		}
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, filters.Limit)
		argIndex++
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	var products []productEntity.Product
	for rows.Next() {

		// id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at, updated_at
		product := productEntity.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Notes, &product.Price, &product.Stock, &product.Location, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt)
		helper.ErrorIfPanic(err)
		products = append(products, product)
	}

	return products

}

func (repository *ProductRepository) FindById(ctx context.Context, tx *sql.Tx, id string) productEntity.Product {
	query := "SELECT id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at, updated_at FROM products WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, id)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	product := productEntity.Product{}
	if rows.Next() {
		rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Notes, &product.Price, &product.Stock, &product.Location, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt)
	}
	return product
}
