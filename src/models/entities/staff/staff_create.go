package staffentity

import "time"

type Staff struct {
	Id          string
	Name        string
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
