package types

type GenericSuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type GenericErrorResponse struct {
	Status string               `json:"status"`
	Error  ErrorMessageResponse `json:"error"`
}

type ErrorMessageResponse struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}
