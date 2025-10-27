package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"paymentservice/internal/lib"
	service "paymentservice/internal/service/interfaces"
)

type Handler struct {
	log      *slog.Logger
	purchase service.PaymentService
}

func New(log *slog.Logger, paymentService service.PaymentService) *Handler {
	return &Handler{
		log:      log,
		purchase: paymentService,
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
func (h *Handler) Purchase(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getKey(r)
	if err != nil {
		h.log.Info("failed to get key", "err", err)
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	price, err := h.purchase.Purchase(userID)
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": fmt.Sprintf("purchase is compeleted, total price:%f", price),
	})
}
