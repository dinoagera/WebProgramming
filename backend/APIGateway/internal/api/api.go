package api

import (
	"apigateway/internal/client"
	"apigateway/internal/config"
	"apigateway/internal/handler"
	"apigateway/internal/middleware/auth"
	"apigateway/internal/service"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type API struct {
	log     *slog.Logger
	cfg     *config.Config
	router  *mux.Router
	service *service.Service
	handler *handler.Handler
}

func InitAPI(log *slog.Logger, cfg *config.Config) *API {
	clientCatalog := client.NewCatalogClient(cfg)
	clientAuth := client.NewAuthClient(cfg)
	clientCart := client.NewCartClient(cfg)
	service := service.New(log, clientCatalog, clientAuth, clientCart)
	handler := handler.New(log, service, service, service)
	return &API{
		log:     log,
		cfg:     cfg,
		router:  mux.NewRouter(),
		service: service,
		handler: handler,
	}
}
func (api *API) StartServer() {
	api.setupRouter()
	server := &http.Server{
		Handler:      api.router,
		Addr:         api.cfg.Address,
		ReadTimeout:  api.cfg.HTTPReadTimeout,
		IdleTimeout:  api.cfg.HTTPidleTimeout,
		WriteTimeout: api.cfg.HTTPidleTimeout,
	}
	api.log.Info("starting server", "address:", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		api.log.Info("server failed", "error:", err)
		os.Exit(1)
	}
}
func (api *API) setupRouter() {
	api.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// API routes
	apiRouter := api.router.PathPrefix("/api").Subrouter()
	// Public routes
	public := apiRouter.PathPrefix("").Subrouter()
	public.HandleFunc("/getcatalog", api.handler.GetCatalog).Methods(http.MethodGet)
	public.HandleFunc("/image/{productID}", api.handler.GetImage).Methods(http.MethodGet)
	// Auth routes (public)
	authRoutes := apiRouter.PathPrefix("").Subrouter()
	authRoutes.HandleFunc("/register", api.handler.Register).Methods(http.MethodPost)
	authRoutes.HandleFunc("/login", api.handler.Login).Methods(http.MethodPost)
	// Protected routes
	protected := apiRouter.PathPrefix("").Subrouter()
	protected.Use(auth.New(api.log, api.cfg.JWTSecret))
	protected.HandleFunc("/getcart", api.handler.GetCart).Methods(http.MethodGet)
	protected.HandleFunc("/additem", api.handler.AddItem).Methods(http.MethodPost)
	protected.HandleFunc("/removeitem", api.handler.RemoveItem).Methods(http.MethodPost)
	protected.HandleFunc("/updateitem", api.handler.UpdateItem).Methods(http.MethodPost)
	protected.HandleFunc("/clearcart", api.handler.ClearCart).Methods(http.MethodGet)
	protected.HandleFunc("/getfavourites", api.handler.GetFavourites).Methods(http.MethodGet)
	protected.HandleFunc("/addfavourite", api.handler.GetFavourites).Methods(http.MethodPost)
}
