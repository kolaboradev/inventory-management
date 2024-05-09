package staffService

import (
	"context"

	staffRequest "github.com/kolaboradev/inventory/src/models/web/request/staff"
	staffResponse "github.com/kolaboradev/inventory/src/models/web/response/staff"
)

type StaffServiceInterface interface {
	Register(ctx context.Context, request staffRequest.StaffCreate) staffResponse.StaffResponse
	Login()
}
