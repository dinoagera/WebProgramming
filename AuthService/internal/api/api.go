package api

import (
	_ "authservice/cmd/docs"
	"authservice/internal/config"
	"authservice/internal/handler"
	"authservice/internal/service"
	"authservice/internal/storage"
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
	db      *storage.Storage
	service *service.Service
	handler *handler.Handler
}

func InitApi(log *slog.Logger, cfg *config.Config) *API {
	storage, err := storage.New(log, cfg.StoragePath)
	if err != nil {
		log.Info("failed to init storage")
		os.Exit(1)
	}
	srv := service.New(log, storage, storage)
	handler := handler.New(log, srv, srv)
	api := &API{
		log:     log,
		cfg:     cfg,
		router:  mux.NewRouter(),
		db:      storage,
		service: srv,
		handler: handler,
	}
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
	public.HandleFunc("/register", api.handler.Register).Methods(http.MethodPost)
	public.HandleFunc("/login", api.handler.Auth).Methods(http.MethodPost)
	api.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
	))
}
