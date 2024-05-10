package staffrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kolaboradev/inventory/src/helper"
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

func (repository *StaffRepository) FindByPhoneNumber(ctx context.Context, tx *sql.Tx, PhoneNumber string) (staffentity.Staff, error) {
	query := "SELECT id, name, phone_number, password, created_at, updated_at FROM staffs WHERE phone_number = $1"

	rows, err := tx.QueryContext(ctx, query, PhoneNumber)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	staff := staffentity.Staff{}

	if rows.Next() {
		err = rows.Scan(&staff.Id, &staff.Name, &staff.PhoneNumber, &staff.Password, &staff.CreatedAt, &staff.UpdatedAt)
		helper.ErrorIfPanic(err)
		return staff, nil
	} else {
		return staff, errors.New("staff not found")
	}
}
