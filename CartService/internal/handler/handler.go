package handler

import (
	"CartService/internal/models"
	service "CartService/internal/service/interfaces"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler struct {
	log     *slog.Logger
	getCart service.GetCart
	addItem service.AddItem
}

func New(log *slog.Logger, getCart service.GetCart, addItem service.AddItem) *Handler {
	return &Handler{
		log:     log,
		getCart: getCart,
		addItem: addItem,
	}
}
func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		h.log.Info("failed to get cart, userId is empty")
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	cart, err := h.getCart.GetCart(userID)
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
	var addItemReq models.AddItemRequest
	if err := json.NewDecoder(r.Body).Decode(&addItemReq); err != nil {
		h.log.Info("failed to decode to model add item req", "err", err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		h.log.Info("failed to add item, userId is empty")
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	err := h.addItem.AddItem(userID, addItemReq)
	if err != nil {
		h.log.Info("failed to add item", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "item is added",
	})
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
