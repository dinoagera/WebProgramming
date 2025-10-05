package handler

import (
	"CartService/internal/models"
	service "CartService/internal/service/interfaces"
	"CartService/lib"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler struct {
	log        *slog.Logger
	getCart    service.GetCart
	addItem    service.AddItem
	removeItem service.RemoveItem
	updateItem service.UpdateItem
	clearCart  service.ClearCart
}

func New(log *slog.Logger, getCart service.GetCart, addItem service.AddItem, removeItem service.RemoveItem, updateItem service.UpdateItem, clearCart service.ClearCart) *Handler {
	return &Handler{
		log:        log,
		getCart:    getCart,
		addItem:    addItem,
		removeItem: removeItem,
		updateItem: updateItem,
		clearCart:  clearCart,
	}
}
func (h *Handler) getKey(r *http.Request) (string, error) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		h.log.Info("failed to get key, userId is empty")
		return "", lib.ErrUserIDIsEmpty
	}
	return userID, nil
}
func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getKey(r)
	if err != nil {
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
		"status": "get cart is successfully",
		"cart":   cart,
	})
}
func (h *Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	var addItemReq models.AddItemRequest
	if err := json.NewDecoder(r.Body).Decode(&addItemReq); err != nil {
		h.log.Info("failed to decode to model add item req", "err", err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	userID, err := h.getKey(r)
	if err != nil {
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	err = h.addItem.AddItem(userID, addItemReq)
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
	var removeItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&removeItem); err != nil {
		h.log.Info("failed to decode to model remove item req", "err", err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	userID, err := h.getKey(r)
	if err != nil {
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	err = h.removeItem.RemoveItem(userID, removeItem)
	if err != nil {
		h.log.Info("failed to remove item", "err", err)
		http.Error(w, "failed to remove item", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "item remove is successfully",
	})
}
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updateItem models.UpdateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&updateItem); err != nil {
		h.log.Info("failed to decode to model update item req", "err", err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	userID, err := h.getKey(r)
	if err != nil {
		h.log.Info("failed to get key", "err", err)
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	err = h.updateItem.UpdateItem(userID, updateItem)
	if err != nil {
		h.log.Info("failed to update item", "err", err)
		http.Error(w, "failed to update item", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "item update is successfully",
	})
}
func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getKey(r)
	if err != nil {
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	err = h.clearCart.ClearCart(userID)
	if err != nil {
		h.log.Info("failed to clear cart", "err", err)
		http.Error(w, "failed to clear cart", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "cart is cleared successfully",
	})
}
