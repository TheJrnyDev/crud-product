package utils

import (
	"github.com/labstack/echo/v4"
)

// Response structures
type SuccessResponse struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code,omitempty"`
}

func ResponseSuccess(c echo.Context, statusCode int, data interface{}, message string) error {
	return c.JSON(statusCode, SuccessResponse{
		Result:  data,
		Message: message,
	})
}

// Error responses
func ResponseError(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ErrorResponse{
		Error: message,
		Code:  statusCode,
	})
}
