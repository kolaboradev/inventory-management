package customerService

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/kolaboradev/inventory/src/exception"
	"github.com/kolaboradev/inventory/src/helper"
	customerEntity "github.com/kolaboradev/inventory/src/models/entities/customer"
	customerRequest "github.com/kolaboradev/inventory/src/models/web/request/customer"
	customerResponse "github.com/kolaboradev/inventory/src/models/web/response/customer"
	customerRepository "github.com/kolaboradev/inventory/src/repositories/customer"
)

type CustomerService struct {
	DB           *sql.DB
	customerRepo customerRepository.CustomerRepositoryInterface
	validator    *validator.Validate
}

func NewCustomerService(db *sql.DB, validator *validator.Validate, cr customerRepository.CustomerRepositoryInterface) CustomerServiceInterface {
	return &CustomerService{
		customerRepo: cr,
		DB:           db,
		validator:    validator,
	}
}

func (service *CustomerService) Create(ctx context.Context, request customerRequest.CustomerCreateRequest) customerResponse.CustomerCreateResonse {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	phoneMustExist := service.customerRepo.FindByPhoneNumber(ctx, tx, request.PhoneNumber)
	if phoneMustExist {
		panic(exception.NewConflictError("phoneNumber is already exists"))
	}

	id := helper.UUIDStr()
	timeNow := helper.TimeISO8601()

	customer := customerEntity.Customer{
		Id:          id,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}

	service.customerRepo.Save(ctx, tx, &customer)

	return customerResponse.CustomerCreateResonse{
		Id:          customer.Id,
		PhoneNumber: customer.PhoneNumber,
		Name:        customer.Name,
	}
}

func (service *CustomerService) FindAll(ctx context.Context, filters customerRequest.CustomerFilter) []customerResponse.CustomerGetResonse {
	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	customers := service.customerRepo.FindAll(ctx, tx, filters)

	var customerResponses []customerResponse.CustomerGetResonse
	for _, value := range customers {
		customer := customerResponse.CustomerGetResonse{
			Id:          value.Id,
			PhoneNumber: value.PhoneNumber,
			Name:        value.Name,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}
		customerResponses = append(customerResponses, customer)
	}
	return customerResponses
}
