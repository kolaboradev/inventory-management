package staffService

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	staffentity "github.com/kolaboradev/inventory/src/models/entities/staff"

	staffRequest "github.com/kolaboradev/inventory/src/models/web/request/staff"
	staffResponse "github.com/kolaboradev/inventory/src/models/web/response/staff"
	staffrepository "github.com/kolaboradev/inventory/src/repositories/staff"
	"golang.org/x/crypto/bcrypt"
)

type StaffService struct {
	StaffRepository staffrepository.StaffRepositoryInterface
	DB              *sql.DB
	validator       *validator.Validate
}

func NewStaffService(staffRepo staffrepository.StaffRepositoryInterface, db *sql.DB, validator validator.Validate) StaffServiceInterface {
	return &StaffService{
		StaffRepository: staffRepo,
		DB:              db,
		validator:       &validator,
	}
}

func (service *StaffService) Register(ctx context.Context, request staffRequest.StaffCreate) staffResponse.StaffResponse {
	err := service.validator.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				panic(errRollback)
			}
			panic(err)
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	// _, err = service.StaffRepository.FindByPhoneNumber(request.PhoneNumber)
	// if err != nil {
	// 	panic(err)
	// }

	saltStr := os.Getenv("BCRYPT_SALT")
	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		panic(err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), salt)
	if err != nil {
		panic(err)
	}

	id := uuid.New()
	timeNowStr := time.Now().Format(time.RFC3339)
	timeNow, err := time.Parse(time.RFC3339, timeNowStr)
	if err != nil {
		panic(err)
	}

	staff := staffentity.Staff{
		Id:          id.String(),
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Password:    string(passwordHash),
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}

	service.StaffRepository.Save(ctx, tx, &staff)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["staffId"] = staff.Id
	claims["exp"] = time.Now().Add(time.Hour * 8)
	claims["phoneNumber"] = staff.PhoneNumber

	secretKeys := []byte(os.Getenv("JWT_SECRET"))

	secretToken, err := token.SignedString(secretKeys)
	if err != nil {
		panic(err)
	}

	return staffResponse.StaffResponse{
		Id:          staff.Id,
		Name:        staff.Name,
		PhoneNumber: staff.PhoneNumber,
		AccessToken: secretToken,
	}

}
func (service *StaffService) Login() {}
