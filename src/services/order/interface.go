package orderService

import (
	"context"

	orderRequest "github.com/kolaboradev/inventory/src/models/web/request/order"
	orderResponse "github.com/kolaboradev/inventory/src/models/web/response/order"
)

type OrderServiceInterface interface {
	Create(ctx context.Context, request orderRequest.OrderCreateRequest) orderResponse.OrderResponse
	FindAll(ctx context.Context, filters orderRequest.OrderGet) []orderResponse.OrderGetResponse
}
