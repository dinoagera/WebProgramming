package api

import (
	"apigateway/internal/client"
	"apigateway/internal/config"
	"apigateway/internal/handler"
	"apigateway/internal/service"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	public := api.router.PathPrefix("/api").Subrouter()
	public.HandleFunc("/getcatalog", api.handler.GetCatalog).Methods(http.MethodGet)
	public.HandleFunc("/image/{productID}", api.handler.GetImage).Methods(http.MethodGet)
	public.HandleFunc("/register", api.handler.Register).Methods(http.MethodPost)
	public.HandleFunc("/login", api.handler.Login).Methods(http.MethodPost)
}
