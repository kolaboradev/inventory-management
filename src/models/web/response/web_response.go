package webResponse

type WebResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
