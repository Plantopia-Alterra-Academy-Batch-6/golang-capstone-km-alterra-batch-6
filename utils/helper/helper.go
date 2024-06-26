package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data interface{} `json:"data"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}