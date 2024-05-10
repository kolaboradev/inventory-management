package productResponse

import "time"

type ProductCreateResponse struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
