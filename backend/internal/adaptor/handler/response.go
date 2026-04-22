package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	CodeSuccess = "1000"
	CodeCreated = "2000"
	CodeBadReq  = "4000"
	CodeUnauth  = "8000"
	CodeConflict = "4001"
	CodeError   = "5000"
)

type MessageResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Time    string `json:"time"`
	Data    any    `json:"data,omitempty"`
}

func respondSuccess(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, MessageResponse{
		Code:    CodeSuccess,
		Message: "success",
		Time:    time.Now().Format(time.RFC3339),
		Data:    data,
	})
}

func respondCreated(c echo.Context) error {
	return c.JSON(http.StatusCreated, MessageResponse{
		Code:    CodeCreated,
		Message: "created",
		Time:    time.Now().Format(time.RFC3339),
	})
}

func respondBadRequest(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, MessageResponse{
		Code:    CodeBadReq,
		Message: msg,
		Time:    time.Now().Format(time.RFC3339),
	})
}

func respondUnauthorized(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, MessageResponse{
		Code:    CodeUnauth,
		Message: "unauthorized",
		Time:    time.Now().Format(time.RFC3339),
	})
}

func respondConflict(c echo.Context, msg string) error {
	return c.JSON(http.StatusConflict, MessageResponse{
		Code:    CodeConflict,
		Message: msg,
		Time:    time.Now().Format(time.RFC3339),
	})
}
