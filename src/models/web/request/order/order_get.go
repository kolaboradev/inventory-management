package orderRequest

type OrderGet struct {
	CustomerId string `json:"customerId"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	CreatedAt  string `json:"createdAt"`
}
