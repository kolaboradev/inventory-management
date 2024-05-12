package orderResponse

import "time"

type OrderDetailResponse struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type OrderGetResponse struct {
	TransactionId  string                `json:"transactionId"`
	CustomerId     string                `json:"customerId"`
	ProductDetails []OrderDetailResponse `json:"productDetails"`
	Paid           int                   `json:"paid"`
	Change         int                   `json:"change"`
	CreatedAt      time.Time             `json:"createdAt"`
}
