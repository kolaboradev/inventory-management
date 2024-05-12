package customerRepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kolaboradev/inventory/src/helper"
	customerEntity "github.com/kolaboradev/inventory/src/models/entities/customer"
	customerRequest "github.com/kolaboradev/inventory/src/models/web/request/customer"
)

type CustomerRepository struct {
}

func NewCustomerRepo() CustomerRepositoryInterface {
	return &CustomerRepository{}
}

func (repository *CustomerRepository) Save(ctx context.Context, tx *sql.Tx, customer *customerEntity.Customer) *customerEntity.Customer {
	query := "INSERT INTO customers (id, name, phone_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := tx.ExecContext(ctx, query, customer.Id, customer.Name, customer.PhoneNumber, customer.CreatedAt, customer.UpdatedAt)
	helper.ErrorIfPanic(err)

	return customer
}

func (repository *CustomerRepository) FindByPhoneNumber(ctx context.Context, tx *sql.Tx, phoneNumber string) bool {
	query := "SELECT phone_number FROM customers WHERE phone_number = $1"

	rows, err := tx.QueryContext(ctx, query, phoneNumber)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}

func (repository *CustomerRepository) FindAll(ctx context.Context, tx *sql.Tx, filters customerRequest.CustomerFilter) []customerEntity.Customer {
	query := "SELECT id, name, phone_number, created_at, updated_at FROM customers WHERE 1=1"

	var args []interface{}
	argIndex := 1

	if filters.PhoneNumber != "" {
		phone := "+" + filters.PhoneNumber
		query += fmt.Sprintf(" AND phone_number LIKE $%d || '%%'", argIndex)
		args = append(args, phone)
		argIndex++
	}
	if filters.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, filters.Name)
		argIndex++
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	var customers []customerEntity.Customer
	for rows.Next() {
		customer := customerEntity.Customer{}
		err = rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.CreatedAt, &customer.UpdatedAt)
		helper.ErrorIfPanic(err)

		customers = append(customers, customer)
	}

	return customers
}

func (repository *CustomerRepository) FindById(ctx context.Context, tx *sql.Tx, id string) bool {
	query := "SELECT id FROM customers WHERE id = $1"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}

}
