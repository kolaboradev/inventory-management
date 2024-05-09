package staffResponse

type StaffResponse struct {
	Id          string `json:"UserId"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	AccessToken string `json:"accessToken"`
}
