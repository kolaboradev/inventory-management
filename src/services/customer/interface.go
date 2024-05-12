package customerService

import (
	"context"

	customerRequest "github.com/kolaboradev/inventory/src/models/web/request/customer"
	customerResponse "github.com/kolaboradev/inventory/src/models/web/response/customer"
)

type CustomerServiceInterface interface {
	Create(ctx context.Context, request customerRequest.CustomerCreateRequest) customerResponse.CustomerCreateResonse
	FindAll(ctx context.Context, filters customerRequest.CustomerFilter) []customerResponse.CustomerGetResonse
}
