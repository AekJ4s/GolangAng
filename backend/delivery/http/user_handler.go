package http

import (
	"encoding/json"
	"net/http"
	"strings"

	domainUseCase "golangbackend/domain/usecase"
	"golangbackend/usecase"
)

type UserHandler struct {
	userUseCase domainUseCase.UserUseCase
}

func NewUserHandler(uc domainUseCase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: uc}
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{"invalid request"})
		return
	}

	if req.Username == "" || req.Password == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{"username and password required"})
		return
	}

	if err := h.userUseCase.Register(req.Username, req.Password); err != nil {
		writeJSON(w, http.StatusConflict, errorResponse{err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{"invalid request"})
		return
	}

	token, err := h.userUseCase.Login(req.Username, req.Password)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, errorResponse{err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, loginResponse{Token: token})
}

func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	username, err := usecase.ValidateToken(tokenString)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, errorResponse{"unauthorized"})
		return
	}

	user, err := h.userUseCase.GetProfile(username)
	if err != nil || user == nil {
		writeJSON(w, http.StatusNotFound, errorResponse{"user not found"})
		return
	}

	writeJSON(w, http.StatusOK, user)
}
