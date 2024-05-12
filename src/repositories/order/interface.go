package orderRepository

import (
	"context"
	"database/sql"

	orderEntity "github.com/kolaboradev/inventory/src/models/entities/order"
	orderRequest "github.com/kolaboradev/inventory/src/models/web/request/order"
)

type OrderRepositoryInterface interface {
	Create(ctx context.Context, tx *sql.Tx, orders orderEntity.OrderCreate)
	FindAll(ctx context.Context, tx *sql.Tx, filters orderRequest.OrderGet) []orderEntity.OrderCreate
}
