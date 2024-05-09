package staffrepository

import (
	"context"
	"database/sql"

	staffentity "github.com/kolaboradev/inventory/src/models/entities/staff"
)

type StaffRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, staff *staffentity.Staff) *staffentity.Staff
	FindByPhoneNumber(PhoneNumber string) (staffentity.Staff, error)
}
