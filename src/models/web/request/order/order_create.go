package orderRequest

// {
// 	"customerId": "", // ID Should be string
// 	"productDetails": [
// 		{
// 			"productId": "",
// 			"quantity": 1 // not null, min: 1
// 		}
// 	], // ID Should be string, minItems: 1
// 	"paid": 1, // not null, min: 1, validate the change based on all product price
// 	"change": 0, // not null, min 0
// }

type OrderProductDetails struct {
	ProductId string `validate:"required"`
	Quantity  int    `validate:"required,min=1"`
}

type OrderCreateRequest struct {
	CustomerId     string                `validate:"required"`
	ProductDetails []OrderProductDetails `validate:"required,min=1,dive"`
	Paid           int                   `validate:"required,min=1"`
	Change         int                   `validate:"required,min=0"`
}
