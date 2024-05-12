package orderEntity

import "time"

type Order struct {
	Id         string
	CustomerId string
	Paid       int
	Change     int
	CreatedAt  time.Time
}

type OrderDetail struct {
	OrderId   string
	ProductId string
	Quantity  int
	CreatedAt time.Time
}
