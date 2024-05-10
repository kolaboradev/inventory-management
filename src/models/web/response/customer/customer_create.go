package customerResponse

type CustomerResponse struct {
	Id          string `json:"UserId"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	AccessToken string `json:"accessToken"`
}
