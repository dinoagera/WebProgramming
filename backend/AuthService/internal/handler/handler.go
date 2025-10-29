package handler

import (
	service "authservice/internal/service/interfaces"
	liberror "authservice/lib/errors"
	"authservice/lib/validator"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type Handler struct {
	log      *slog.Logger
	login    service.Login
	register service.Register
}
type RegisterFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New(log *slog.Logger, login service.Login, register service.Register) *Handler {
	return &Handler{
		log:      log,
		login:    login,
		register: register,
	}
}

// @Summary SignUp
// @Tags CreateUser
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body RegisterFormat true "account info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.log.Info("method is allowed")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req RegisterFormat
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Info("decode to failed in register handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := validator.ValidateEmail(req.Email)
	switch {
	case errors.Is(err, liberror.ErrEmailEmpty):
		h.log.Info("failed to register", "err", err)
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	case errors.Is(err, liberror.ErrEmailNotAllowed):
		h.log.Info("failed to register", "err", err)
		http.Error(w, "email is not allowed", http.StatusBadRequest)
		return
	}
	err = validator.ValidatePassword(req.Password)
	if err != nil {
		http.Error(w, "Password less 6 symbol", http.StatusBadRequest)
		return
	}
	err = h.register.Register(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, liberror.ErrEmailBusy) {
			h.log.Info("failed to register", "err", err)
			http.Error(w, "email is busy", http.StatusBadRequest)
			return
		}
		h.log.Info("failed to register", "err:", err)
		http.Error(w, "failed to register", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully"})
}

// @Summary SignIn
// @Tags Login
// @Description login account
// @ID login-account
// @Accept json
// @Produce json
// @Param input body RegisterFormat true "account info"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /login [post]
func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req RegisterFormat
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Info("decode to failed in login handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := validator.ValidateEmail(req.Email)
	switch {
	case errors.Is(err, liberror.ErrEmailEmpty):
		h.log.Info("failed to register", "err", err)
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	case errors.Is(err, liberror.ErrEmailNotAllowed):
		h.log.Info("failed to register", "err", err)
		http.Error(w, "email is not allowed", http.StatusBadRequest)
		return
	}
	err = validator.ValidatePassword(req.Password)
	if err != nil {
		http.Error(w, "Password less 6 symbol", http.StatusBadRequest)
		return
	}
	token, err := h.login.Login(req.Email, req.Password)
	if err != nil {
		h.log.Info("failed to login", "err:", err)
		http.Error(w, "failed to login", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
