package helper

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Field   string `json:"field"`
	Error   string `json:"error"`
}
