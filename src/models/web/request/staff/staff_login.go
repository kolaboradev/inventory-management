package staffRequest

type StaffLogin struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,phone_number,min=10,max=16"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}
