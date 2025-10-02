package handler

import (
	service "CartService/internal/service/interfaces"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	log     *slog.Logger
	getCart service.GetCart
}

func New(log *slog.Logger, getCart service.GetCart) *Handler {
	return &Handler{
		log:     log,
		getCart: getCart,
	}
}
func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		h.log.Info("failed to get cart, userId is empty")
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		h.log.Info("failed to convert userid", "err", err)
		http.Error(w, "failed to get cart", http.StatusInternalServerError)
		return
	}
	cart, err := h.getCart.GetCart(id)
	if err != nil {
		h.log.Info("failed to get cart", "err", err)
		http.Error(w, "failed to get cart", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "successfully",
		"cart":   cart,
	})
}
func (h *Handler) AddItem(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	//
}
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	//
}
func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	//
}
