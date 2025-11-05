package handler

import (
	service "catalogservice/internal/service/interfaces"
	"catalogservice/lib"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

type Handler struct {
	log           *slog.Logger
	getCatalog    service.GetCatalog
	getImage      service.GetImage
	getFavourites service.GetFavourites
}

func New(log *slog.Logger, getCatalog service.GetCatalog, getImage service.GetImage, getFavourites service.GetFavourites) *Handler {
	return &Handler{
		log:           log,
		getCatalog:    getCatalog,
		getImage:      getImage,
		getFavourites: getFavourites,
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
func (h *Handler) GetCatalog(w http.ResponseWriter, r *http.Request) {
	goods, err := h.getCatalog.GetCatalog()
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
	imageData, err := h.getImage.GetImage(productID)
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
func (h *Handler) GetFavourites(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getKey(r)
	if err != nil {
		h.log.Info("failed to get key", "err", err)
		http.Error(w, "not authorization", http.StatusUnauthorized)
		return
	}
	favourites, err := h.getFavourites.GetFavourites(userID)
	if err != nil {
		h.log.Info("failed to get favourites", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "get catalog is successfully",
		"favourites": favourites,
	})
}
