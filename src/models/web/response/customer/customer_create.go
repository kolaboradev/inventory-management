package customerResponse

type CustomerCreateResonse struct {
	Id          string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}
