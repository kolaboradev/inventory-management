package customerentity

import "time"

type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
