package orderService

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/kolaboradev/inventory/src/exception"
	"github.com/kolaboradev/inventory/src/helper"
	orderEntity "github.com/kolaboradev/inventory/src/models/entities/order"
	orderRequest "github.com/kolaboradev/inventory/src/models/web/request/order"
	orderResponse "github.com/kolaboradev/inventory/src/models/web/response/order"
	customerRepository "github.com/kolaboradev/inventory/src/repositories/customer"
	orderRepository "github.com/kolaboradev/inventory/src/repositories/order"
	productRepository "github.com/kolaboradev/inventory/src/repositories/product"
)

type OrderService struct {
	DB           *sql.DB
	validator    *validator.Validate
	orderRepo    orderRepository.OrderRepositoryInterface
	productRepo  productRepository.ProductRepositoryInterface
	customerRepo customerRepository.CustomerRepositoryInterface
}

func NewOrderService(db *sql.DB, validator *validator.Validate, orderRepo orderRepository.OrderRepositoryInterface, productRepo productRepository.ProductRepositoryInterface, customerRepo customerRepository.CustomerRepositoryInterface) OrderServiceInterface {
	return &OrderService{
		DB:           db,
		validator:    validator,
		orderRepo:    orderRepo,
		productRepo:  productRepo,
		customerRepo: customerRepo,
	}
}

func (service *OrderService) Create(ctx context.Context, request orderRequest.OrderCreateRequest) orderResponse.OrderResponse {
	err := service.validator.Struct(request)
	helper.ErrorIfPanic(err)

	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	isCustomer := service.customerRepo.FindById(ctx, tx, request.CustomerId)
	if !isCustomer {
		panic(exception.NewNotFoundError("customer not found"))
	}

	pricePaid := 0

	for _, value := range request.ProductDetails {
		isValid := service.productRepo.FindByIdBool(ctx, tx, value.ProductId)
		if !isValid {
			panic(exception.NewNotFoundError("product is not found"))
		}
		product := service.productRepo.FindById(ctx, tx, value.ProductId)
		if !product.IsAvailable {
			panic(exception.NewBadRequestError("product not available"))
		}
		pricePaid += product.Price * value.Quantity
	}

	if pricePaid > request.Paid {
		panic(exception.NewBadRequestError("is not enough based on all bought product"))
	}
	changeReturn := request.Paid - pricePaid
	if changeReturn != *request.Change {
		panic(exception.NewBadRequestError("is not right, based on all bought product, and what is paid"))
	}

	idOrder := helper.UUIDStr()
	timeNow := helper.TimeISO8601()

	orders := orderEntity.OrderCreate{
		Order: orderEntity.Order{
			Id:         idOrder,
			CustomerId: request.CustomerId,
			Paid:       request.Paid,
			Change:     *request.Change,
			CreatedAt:  timeNow,
		},
	}

	var orderSlice []orderEntity.OrderDetail

	for _, order := range request.ProductDetails {
		orderDetail := orderEntity.OrderDetail{
			OrderId:   orders.Order.Id,
			ProductId: order.ProductId,
			Quantity:  order.Quantity,
			CreatedAt: timeNow,
		}
		orderSlice = append(orderSlice, orderDetail)
	}

	orders.OrderDetail = orderSlice

	service.orderRepo.Create(ctx, tx, orders)

	return orderResponse.OrderResponse{
		Paid: changeReturn,
	}

}

func (service *OrderService) FindAll(ctx context.Context, filters orderRequest.OrderGet) []orderResponse.OrderGetResponse {
	tx, err := service.DB.Begin()
	helper.ErrorIfPanic(err)
	defer helper.RollbackOrCommit(tx)

	orders := service.orderRepo.FindAll(ctx, tx, filters)

	var orderResponses []orderResponse.OrderGetResponse
	for _, order := range orders {

		var odResponse []orderResponse.OrderDetailResponse
		for _, value := range order.OrderDetail {
			odResponse = append(odResponse, orderResponse.OrderDetailResponse{
				ProductId: value.ProductId,
				Quantity:  value.Quantity,
			})
		}

		orderResponses = append(orderResponses, orderResponse.OrderGetResponse{
			TransactionId:  order.Order.Id,
			CustomerId:     order.Order.CustomerId,
			ProductDetails: odResponse,
			Paid:           order.Order.Paid,
			Change:         order.Order.Change,
			CreatedAt:      order.Order.CreatedAt,
		})
	}
	return orderResponses
}
