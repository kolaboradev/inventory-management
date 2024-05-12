package customerResponse

import "time"

type CustomerGetResonse struct {
	Id          string    `json:"userId"`
	PhoneNumber string    `json:"phoneNumber"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
