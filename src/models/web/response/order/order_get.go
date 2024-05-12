package orderResponse

type OrderDetailResponse struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type OrderGetResponse struct {
	TransactionId  string `json:"transactionId"`
	CustomerId     string `json:"customerId"`
	ProductDetails []OrderDetailResponse
	Paid           int `json:"Paid"`
	Change         int `json:"change"`
}
