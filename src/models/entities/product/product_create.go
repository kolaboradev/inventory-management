package productEntity

import "time"

type Product struct {
	Id          string
	Name        string
	Sku         string
	Category    string
	ImageUrl    string
	Notes       string
	Price       int
	Stock       int
	Location    string
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
