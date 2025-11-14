package api

import (
	"log/slog"
	"net/http"
	"os"
	"paymentservice/internal/client"
	"paymentservice/internal/config"
	"paymentservice/internal/handler"
	"paymentservice/internal/service"

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
	client := client.NewClient(cfg)
	srv := service.New(log, client)
	handler := handler.New(log, srv)
	api := &API{
		log:     log,
		cfg:     cfg,
		router:  mux.NewRouter(),
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
		Addr:         api.cfg.Address,
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
	public.HandleFunc("/purchase", api.handler.Purchase).Methods(http.MethodGet)
}
