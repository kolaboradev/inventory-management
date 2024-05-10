package staffService

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/kolaboradev/inventory/src/exception"
	"github.com/kolaboradev/inventory/src/helper"
	staffentity "github.com/kolaboradev/inventory/src/models/entities/staff"

	staffRequest "github.com/kolaboradev/inventory/src/models/web/request/staff"
	staffResponse "github.com/kolaboradev/inventory/src/models/web/response/staff"
	staffrepository "github.com/kolaboradev/inventory/src/repositories/staff"
)

type StaffService struct {
	StaffRepository staffrepository.StaffRepositoryInterface
	DB              *sql.DB
	validator       *validator.Validate
}

func NewStaffService(staffRepo staffrepository.StaffRepositoryInterface, db *sql.DB, validator *validator.Validate) StaffServiceInterface {
	return &StaffService{
		StaffRepository: staffRepo,
		DB:              db,
		validator:       validator,
	}
}

func (service *StaffService) Register(ctx context.Context, request staffRequest.StaffCreate) staffResponse.StaffResponse {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)

	defer helper.RollbackOrCommit(tx)

	_, err = service.StaffRepository.FindByPhoneNumber(ctx, tx, request.PhoneNumber)
	if err == nil {
		panic(exception.NewConflictError("Phone number exists"))
	}

	hashPassword := helper.HashPassword(request.Password)

	id := helper.UUIDStr()
	timeNow := helper.TimeISO8601()

	staff := staffentity.Staff{
		Id:          id,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Password:    hashPassword,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}

	service.StaffRepository.Save(ctx, tx, &staff)

	token := helper.GenerateTokenJWT(staff)

	return staffResponse.StaffResponse{
		Id:          staff.Id,
		Name:        staff.Name,
		PhoneNumber: staff.PhoneNumber,
		AccessToken: token,
	}

}
func (service *StaffService) Login(ctx context.Context, request staffRequest.StaffLogin) staffResponse.StaffResponse {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)

	defer helper.RollbackOrCommit(tx)

	staff, err := service.StaffRepository.FindByPhoneNumber(ctx, tx, request.PhoneNumber)
	if err != nil {
		panic(exception.NewNotFoundError("staff is not found"))
	}

	passwordIsValid := helper.CompareHashPassword(staff.Password, request.Password)

	if !passwordIsValid {
		panic(exception.NewBadRequestError("password is wrong"))
	}

	token := helper.GenerateTokenJWT(staff)

	return staffResponse.StaffResponse{
		Id:          staff.Id,
		Name:        staff.Name,
		PhoneNumber: staff.PhoneNumber,
		AccessToken: token,
	}
}
