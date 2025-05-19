package responseHelper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Data        any    `json:"data"`
	Error       any    `json:"error"`
	Validations []string `json:"validations"`
}

func SuccessResponse(data any, message string) *Response {
	return &Response{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse( err error, message string, data any, code int) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Error:   err.Error(),
		Data:    data,
	}
}

func ValidationErrorResponse(err error, message string, data any, code int) *Response {
	var validationErrors []string
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, err.Error())
	}
	return &Response{
		Code:        code,
		Message:     message,
		Validations: validationErrors,
		Data:        data,
	}
}
