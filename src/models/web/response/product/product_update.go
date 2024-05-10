package productResponse

import "time"

type ProductUpdateResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Sku         string    `json:"sku"`
	Category    string    `json:"category"`
	ImageUrl    string    `json:"imageUrl"`
	Notes       string    `json:"notes"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"isAvailable"`
	UpdateAt    time.Time `json:"createdAt"`
}
