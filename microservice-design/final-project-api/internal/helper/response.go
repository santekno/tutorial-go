package helper

import (
	"strings"
)

// Represents a response from the API.
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Empty Object is used when data doesn't want to be null on json
type EmptyObj struct{}

// BuildSuccessResponse is used to build a success response.
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

// BuildErrorResponse is used to build an error response.
func BuildErrorResponse(message string, errors string, data interface{}) Response {
	splittedError := strings.Split(errors, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}
