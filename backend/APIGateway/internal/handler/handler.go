package handler

import (
	service "apigateway/internal/service/interfaces"
	"apigateway/lib"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
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
		if err == lib.ErrCatalogIsEmpty {
			h.log.Info("catalog is empty", "err", err)
			http.Error(w, "Catalog is empty", http.StatusInternalServerError)
			return
		}
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
	productID := strings.TrimPrefix(r.URL.Path, "/api/image/")
	if productID == "" {
		h.log.Info("request have not product id")
		http.Error(w, "Product ID required", http.StatusBadRequest)
		return
	}
	imageData, err := h.catalogService.GetImage(productID)
	if err != nil {
		if err == lib.ErrImageNotFound {
			h.log.Info("image not found")
			http.Error(w, "Image not found", http.StatusInternalServerError)
			return
		}
		h.log.Info("failed to get image", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(imageData)
}
