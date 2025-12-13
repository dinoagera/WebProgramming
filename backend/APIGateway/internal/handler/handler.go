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

type CatalogResponse struct {
	Status  string        `json:"status"`
	Catalog []models.Good `json:"catalog"`
}
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

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body models.AuthRequest true "account info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
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

// @Summary SignIn
// @Tags auth
// @Description login account
// @ID login-account
// @Accept json
// @Produce json
// @Param input body models.AuthRequest true "account info"
// @Success 200 {object} map[string]string "Return JWT token"
// @Failure 400 {object} map[string]string "Validation error"
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Info("decode to failed in login handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		h.log.Info("failed to login", "err:", err)
		http.Error(w, fmt.Sprintf("error:%s", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

// @Summary GetAllCatalog
// @Tags catalog
// @Description Return all goods in catalog
// @ID get-all-catalog
// @Produce json
// @Success 200 {array} models.CatalogResponse
// @Failure 400 {object} map[string]string
// @Router /getcatalog [get]
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
	json.NewEncoder(w).Encode(models.CatalogResponse{
		Status:  "get catalog is successfully",
		Catalog: goods,
	})
}

// @Summary GetProductByID
// @Tags catalog
// @Description Return good by id
// @ID get-product
// @Produce json
// @Success 200 {array} models.ProductResponse
// @Failure 400 {object} map[string]string
// @Router /product/{productID} [get]
func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := strings.TrimPrefix(r.URL.Path, "/api/product/")
	good, err := h.catalogService.GetProduct(productID)
	if err != nil {
		h.log.Info("failed to get product", "err", err)
		http.Error(w, "Internal server", http.StatusInternalServerError)
		return
	}
	h.log.Info("get product is successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ProductResponse{
		Status:  "get product is successfully",
		Product: good,
	})
}

// @Summary GetFavouritesGoods
// @Tags catalog
// @Description Return all favourites goods by user
// @ID get-favourites-goods
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.FavouritesResponse
// @Failure 400 {object} map[string]string
// @Router /getfavourites [get]
func (h *Handler) GetFavourites(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.log.Info("faield to get userID")
		http.Error(w, "Unauthorization", http.StatusUnauthorized)
		return
	}
	favourites, err := h.catalogService.GetFavourites(userID)
	if err != nil {
		h.log.Info("failed to get favourites", "err", err)
		http.Error(w, "failed to get favourites", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.FavouritesResponse{
		Status:     "get favourites is successfully",
		Favourites: favourites,
	})
}

// @Summary Add Favourite
// @Tags catalog
// @Description Add item to favourites
// @ID add-favourite
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.AddFavouriteRequest true "productID info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /addfavourite [post]
func (h *Handler) AddFavourite(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.log.Info("failed to get userID")
		http.Error(w, "Unauthorization", http.StatusUnauthorized)
		return
	}
	var addFavouriteRequest models.AddFavouriteRequest
	if err := json.NewDecoder(r.Body).Decode(&addFavouriteRequest); err != nil {
		h.log.Info("decode to failed in add favourtie item handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.catalogService.AddFavourite(userID, addFavouriteRequest.ProductID)
	if err != nil {
		h.log.Info("failed to add favourites", "err", err)
		http.Error(w, "failed to add favourites", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "favourite item is added successfully",
	})
}

// @Summary Remove Favourite
// @Tags catalog
// @Description Remove item from favourites items
// @ID remove-favourite
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.RemoveFavouriteRequest true "productID info"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /removefavourite [post]
func (h *Handler) RemoveFavourite(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.log.Info("failed to get userID")
		http.Error(w, "Unauthorization", http.StatusUnauthorized)
		return
	}
	var removeFavouriteRequest models.RemoveFavouriteRequest
	if err := json.NewDecoder(r.Body).Decode(&removeFavouriteRequest); err != nil {
		h.log.Info("decode to failed in add favourtie item handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.catalogService.RemoveFavourite(userID, removeFavouriteRequest.ProductID)
	if err != nil {
		h.log.Info("failed to removed favourites", "err", err)
		http.Error(w, "failed to removed favourites", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "favourite item is removed successfully",
	})
}

// @Summary Get image
// @Tags catalog
// @Description Return product image by ID
// @ID get-image
// @Produce image/png
// @Param productid path string true "productID"
// @Success 200 {file} byte "Product image"
// @Failure 400 {object} map[string]string
// @Router /image/{productID} [get]
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
	contentType := http.DetectContentType(imageData)
	w.Header().Set("Content-Type", contentType)
	w.Write(imageData)
}

// @Summary Get cart
// @Tags cart
// @Description Return user cart by user_id from JWT token
// @ID get-cart
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.CartResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /getcart [get]
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
	json.NewEncoder(w).Encode(models.CartResponse{
		Status: "get cart is successfully",
		Cart:   cart,
	})
}

// @Summary Add Item
// @Tags cart
// @Description Add item to cart
// @ID add-item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.AddItemRequest true "item info"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /additem [post]
func (h *Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.log.Info("failed to get userID")
		http.Error(w, "Unauthorization", http.StatusUnauthorized)
		return
	}
	var addItemRequest models.AddItemRequest
	if err := json.NewDecoder(r.Body).Decode(&addItemRequest); err != nil {
		h.log.Info("decode to failed in add item handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.cartService.AddItem(userID, addItemRequest.ProductID, addItemRequest.Quantity, addItemRequest.Price, addItemRequest.Category)
	if err != nil {
		h.log.Info("failed to add item", "err:", err)
		http.Error(w, "Failed to add item", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "item is added",
	})
}

// @Summary Remove Item
// @Tags cart
// @Description Remove item from cart by product_id
// @ID remove-item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.RemoveItemRequest true "product id"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /removeitem [post]
func (h *Handler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.log.Info("failed to get userID")
		http.Error(w, "Unauthorization", http.StatusUnauthorized)
		return
	}
	var removeItemRequest models.RemoveItemRequest
	if err := json.NewDecoder(r.Body).Decode(&removeItemRequest); err != nil {
		h.log.Info("decode to failed in remove item handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.cartService.RemoveItem(userID, removeItemRequest.ProductID)
	if err != nil {
		h.log.Info("failed to remove item", "err:", err)
		http.Error(w, "Failed to remove item", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "item is removed",
	})
}

// @Summary Update Item
// @Tags cart
// @Description Update item up or down quantity, 0 - down , >0 - up
// @ID update-item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.UpdateItemRequest true "product id"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /updateitem [post]
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.log.Info("failed to get userID")
		http.Error(w, "Unauthorization", http.StatusUnauthorized)
		return
	}
	var updateItemRequest models.UpdateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&updateItemRequest); err != nil {
		h.log.Info("decode to failed in update item handler", "err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.cartService.UpdateItem(userID, updateItemRequest.ProductID, updateItemRequest.TypeOperation)
	if err != nil {
		h.log.Info("failed to update item", "err:", err)
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "item is updated",
	})
}

// @Summary Clear Cart
// @Tags cart
// @Description Clear cart by user id from JWT token
// @ID clear-cart
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clearcart [delete]
func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.log.Info("failed to get userID")
		http.Error(w, "Unauthorization", http.StatusUnauthorized)
		return
	}
	err := h.cartService.ClearCart(userID)
	if err != nil {
		h.log.Info("failed to clear cart", "err:", err)
		http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "cart is cleared",
	})
}

func (h *Handler) GetMale(w http.ResponseWriter, r *http.Request) {
	goods, err := h.catalogService.GetMale()
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
	json.NewEncoder(w).Encode(models.CatalogResponse{
		Status:  "get catalog is successfully",
		Catalog: goods,
	})
}

func (h *Handler) GetFemale(w http.ResponseWriter, r *http.Request) {
	goods, err := h.catalogService.GetFemale()
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
	json.NewEncoder(w).Encode(models.CatalogResponse{
		Status:  "get catalog is successfully",
		Catalog: goods,
	})
}
