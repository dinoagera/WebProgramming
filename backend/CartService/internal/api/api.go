package api

import (
	"CartService/internal/config"
	"CartService/internal/handler"
	"CartService/internal/service"
	"CartService/internal/storage"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type API struct {
	log     *slog.Logger
	cfg     *config.Config
	router  *mux.Router
	db      *storage.Storage
	service *service.Service
	handler *handler.Handler
}

func InitApi(log *slog.Logger, cfg *config.Config, client *redis.Client) *API {
	storage := storage.New(log, client, cfg)
	srv := service.New(log, storage, storage, storage, storage)
	handler := handler.New(log, srv, srv, srv, srv)
	api := &API{
		log:     log,
		cfg:     cfg,
		router:  mux.NewRouter(),
		db:      storage,
		service: srv,
		handler: handler,
	}
	log.Info("API initialized successfully")
	return api
}
func (api *API) StartServer() {
	api.setupRouter()
	server := &http.Server{
		Handler:      api.router,
		Addr:         api.cfg.ServerAddr,
		ReadTimeout:  api.cfg.ReadTimeout,
		WriteTimeout: api.cfg.IdleTimeout,
		IdleTimeout:  api.cfg.IdleTimeout,
	}
	api.log.Info("starting server", "address:", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		api.log.Info("server failed", "error:", err)
		os.Exit(1)
	}
}
func (api *API) setupRouter() {
	public := api.router.PathPrefix("/api").Subrouter()
	public.HandleFunc("/getCart", api.handler.GetCart).Methods(http.MethodGet)
}
