package handler

import (
	"apigateway/internal/middleware/auth"
	"apigateway/internal/models"
	service "apigateway/internal/service/interfaces"
	"apigateway/lib"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

type Handler struct {
	log            *slog.Logger
	catalogService service.CatalogService
	authService    service.AuthService
	cartService    service.CartService
}

func New(log *slog.Logger, catalogService service.CatalogService, authService service.AuthService, cartService service.CartService) *Handler {
	return &Handler{
		log:            log,
		catalogService: catalogService,
		authService:    authService,
		cartService:    cartService,
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
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Info("decode to failed in register handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.authService.Register(req.Email, req.Password)
	if err != nil {
		h.log.Info("failed to register", "err:", err)
		http.Error(w, fmt.Sprintf("error:%s", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully"})
}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Info("decode to failed in login handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		h.log.Info("failed to register", "err:", err)
		http.Error(w, fmt.Sprintf("error:%s", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.log.Info("faield to get userID")
		http.Error(w, "Unauthorization", http.StatusUnauthorized)
		return
	}
	cart, err := h.cartService.GetCart(userID)
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
