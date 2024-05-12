package orderRepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kolaboradev/inventory/src/helper"
	orderEntity "github.com/kolaboradev/inventory/src/models/entities/order"
	orderRequest "github.com/kolaboradev/inventory/src/models/web/request/order"
)

type OrderRepository struct{}

func NewOrderRepo() OrderRepositoryInterface {
	return &OrderRepository{}
}

func (repository *OrderRepository) Create(ctx context.Context, tx *sql.Tx, orders orderEntity.OrderCreate) {
	query := "INSERT INTO orders (id, customer_id, paid, change, created_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := tx.ExecContext(ctx, query, orders.Order.Id, orders.Order.CustomerId, orders.Order.Paid, orders.Order.Change, orders.Order.CreatedAt)
	helper.ErrorIfPanic(err)

	qPrepare1 := "INSERT INTO order_details (order_id, product_id, quantity, created_at) VALUES ($1, $2, $3, $4)"

	stm, err := tx.PrepareContext(ctx, qPrepare1)
	helper.ErrorIfPanic(err)
	defer stm.Close()

	qPrepare2 := "UPDATE products SET stock = stock - $1 WHERE id = $2"
	stm2, err := tx.PrepareContext(ctx, qPrepare2)
	helper.ErrorIfPanic(err)
	defer stm2.Close()

	for _, order := range orders.OrderDetail {
		_, err = stm.ExecContext(ctx, orders.Order.Id, order.ProductId, order.Quantity, order.CreatedAt)
		helper.ErrorIfPanic(err)
	}

	for _, order := range orders.OrderDetail {
		_, err = stm2.ExecContext(ctx, order.Quantity, order.ProductId)
		helper.ErrorIfPanic(err)
	}

}

func (repository *OrderRepository) FindAll(ctx context.Context, tx *sql.Tx, filters orderRequest.OrderGet) []orderEntity.OrderCreate {
	query := "SELECT id, customer_id, paid, change, created_at FROM orders WHERE 1=1"

	var args []interface{}
	argIndex := 1

	if filters.CustomerId != "" {
		query += fmt.Sprintf(" AND customer_id = $%d", argIndex)
		args = append(args, filters.CustomerId)
		argIndex++
	}
	if filters.CreatedAt != "" {
		if filters.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		} else if filters.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		} else {
			query += " ORDER BY created_at desc"
		}
	}

	if filters.CreatedAt == "" {
		query += " ORDER BY created_at desc"
	}

	if filters.Limit >= 0 {
		if filters.Limit == 0 {
			filters.Limit = 5
		}
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, filters.Limit)
		argIndex++
	}
	if filters.Offset >= 0 {
		filters.Offset += 1
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, filters.Offset)
		argIndex++
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	var orders []orderEntity.Order
	for rows.Next() {
		order := orderEntity.Order{}
		err = rows.Scan(&order.Id, &order.CustomerId, &order.Paid, &order.Change, &order.CreatedAt)
		helper.ErrorIfPanic(err)

		orders = append(orders, order)
	}

	queryPrepare := "SELECT product_id, quantity FROM order_details WHERE order_id = $1"

	stm, err := tx.PrepareContext(ctx, queryPrepare)
	helper.ErrorIfPanic(err)
	defer stm.Close()

	var ordersCreate []orderEntity.OrderCreate
	for _, value := range orders {
		rowsStm, err := stm.QueryContext(ctx, value.Id)
		helper.ErrorIfPanic(err)
		defer rowsStm.Close()

		var orderDetailSlice []orderEntity.OrderDetail
		for rowsStm.Next() {
			orderD := orderEntity.OrderDetail{}
			rowsStm.Scan(&orderD.ProductId, &orderD.Quantity)
			orderDetailSlice = append(orderDetailSlice, orderD)
		}
		ordersCreate = append(ordersCreate, orderEntity.OrderCreate{
			Order:       value,
			OrderDetail: orderDetailSlice,
		})
	}
	return ordersCreate
}
