package customerRepository

import (
	"context"
	"database/sql"

	customerEntity "github.com/kolaboradev/inventory/src/models/entities/customer"
	customerRequest "github.com/kolaboradev/inventory/src/models/web/request/customer"
)

type CustomerRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, customer *customerEntity.Customer) *customerEntity.Customer
	FindByPhoneNumber(ctx context.Context, tx *sql.Tx, phoneNumber string) bool
	FindById(ctx context.Context, tx *sql.Tx, id string) bool
	FindAll(ctx context.Context, tx *sql.Tx, filters customerRequest.CustomerFilter) []customerEntity.Customer
}
