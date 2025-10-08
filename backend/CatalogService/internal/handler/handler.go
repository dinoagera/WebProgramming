package handler

import (
	service "catalogservice/internal/service/interfaces"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler struct {
	log        *slog.Logger
	getCatalog service.GetCatalog
}

func New(log *slog.Logger) *Handler {
	return &Handler{log: log}
}
func (h *Handler) GetCatalog(w http.ResponseWriter, r *http.Request) {
	goods, err := h.getCatalog.GetCatalog()
	if err != nil {
		h.log.Info("failed to get catalog", "err", err)
		http.Error(w, "Internal server", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "get catalog is successfully",
		"catalog": goods,
	})
}
