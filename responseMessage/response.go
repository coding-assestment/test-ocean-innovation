package responseMessage

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Responses struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int64       `json:"code"`
}
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int64       `json:"code"`
}

func Error(c echo.Context, data interface{}, err error) error {
	return c.JSON(http.StatusInternalServerError, Response{Message: err.Error(), Data: data, Code: http.StatusInternalServerError})
}

func NewNotFound(c echo.Context, data interface{}, message string) error {

	return c.JSON(http.StatusNotFound, Response{Message: message, Data: data, Code: http.StatusNotFound})
}

func NewSuccess(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusCreated, Response{Message: message, Data: data, Code: http.StatusCreated})
}

func ListSuccess(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusOK, Responses{Message: message, Data: data, Code: http.StatusOK})
}
