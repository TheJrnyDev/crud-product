package utils

import (
	"github.com/labstack/echo/v4"
)

// Response structures following JSend
type JsendResponse struct {
	Status  string      `json:"status,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(c echo.Context, code int, data interface{}, message string) error {
	return c.JSON(code, JsendResponse{
		Status:  "success",
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func ResponseFail(c echo.Context, code int, data interface{}, message string) error {
	return c.JSON(code, JsendResponse{
		Status:  "success",
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func ResponseError(c echo.Context, code int, data interface{}, message string) error {
	return c.JSON(code, JsendResponse{
		Status:  "error",
		Code:    code,
		Data:    data,
		Message: message,
	})
}
