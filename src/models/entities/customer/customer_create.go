package customerEntity

import "time"

type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
