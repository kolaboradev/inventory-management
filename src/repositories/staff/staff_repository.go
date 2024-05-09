package staffrepository

import (
	"context"
	"database/sql"

	staffentity "github.com/kolaboradev/inventory/src/models/entities/staff"
)

type StaffRepository struct {
}

func NewStaffRepository() StaffRepositoryInterface {
	return &StaffRepository{}
}

func (repository *StaffRepository) Save(ctx context.Context, tx *sql.Tx, staff *staffentity.Staff) *staffentity.Staff {
	query := "INSERT INTO staffs (id, name, phone_number, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := tx.ExecContext(ctx, query, staff.Id, staff.Name, staff.PhoneNumber, staff.Password, staff.CreatedAt, staff.UpdatedAt)

	if err != nil {
		panic(err)
	}

	return staff
}

func (repository *StaffRepository) FindByPhoneNumber(PhoneNumber string) (staffentity.Staff, error) {
	panic("implement me")
}
