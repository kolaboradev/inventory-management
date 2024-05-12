package customerRequest

type CustomerCreateRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,phone_number,min=10,max=16"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
}
