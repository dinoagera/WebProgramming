package handler

import (
	service "apigateway/internal/service/interfaces"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler struct {
	log            *slog.Logger
	catalogService service.CatalogService
}

func New(log *slog.Logger, catalogService service.CatalogService) *Handler {
	return &Handler{
		log:            log,
		catalogService: catalogService,
	}
}
func (h *Handler) GetCatalog(w http.ResponseWriter, r *http.Request) {
	goods, err := h.catalogService.GetCatalog()
	if err != nil {
		h.log.Info("failed to get catalog", "err", err)
		http.Error(w, "Internal server", http.StatusInternalServerError)
		return
	}
	h.log.Info("get catalog is successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "get catalog is successfully",
		"catalog": goods,
	})
}
func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it s empty handler"))
}
