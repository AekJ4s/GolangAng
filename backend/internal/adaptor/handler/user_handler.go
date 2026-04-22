package handler

import (
	"net/http"
	"strings"

	"golangbackend/internal/core/port/driving"
	"golangbackend/internal/core/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	svc driving.UserService
}

func NewUserHandler(svc driving.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(c echo.Context) error {
	var req registerRequest
	if err := c.Bind(&req); err != nil {
		return respondBadRequest(c, "invalid request")
	}

	if req.Username == "" || req.Password == "" {
		return respondBadRequest(c, "username and password required")
	}

	if err := h.svc.Register(req.Username, req.Password); err != nil {
		return respondConflict(c, err.Error())
	}

	return respondCreated(c)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return respondBadRequest(c, "invalid request")
	}

	token, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		return respondUnauthorized(c)
	}

	return respondSuccess(c, map[string]string{"token": token})
}

func (h *UserHandler) Profile(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	username, err := service.ValidateToken(tokenString)
	if err != nil {
		return respondUnauthorized(c)
	}

	user, err := h.svc.GetProfile(username)
	if err != nil || user == nil {
		return c.JSON(http.StatusNotFound, MessageResponse{
			Code:    "4004",
			Message: "user not found",
		})
	}

	return respondSuccess(c, map[string]any{
		"id":       user.ID,
		"username": user.Username,
	})
}
