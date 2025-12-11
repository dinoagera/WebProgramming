package api

import (
	"catalogservice/internal/config"
	"catalogservice/internal/handler"
	"catalogservice/internal/service"
	"catalogservice/internal/storage"
	"log/slog"
	"net/http"
	"os"

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

func InitAPI(log *slog.Logger, cfg *config.Config) *API {
	db, err := storage.New(log, cfg.StoragePath, cfg.ImageBasePath)
	if err != nil {
		log.Info("failed to init storage", "err", err)
		os.Exit(1)
	}
	srv := service.New(log, db, db, db, db, db, db)
	handler := handler.New(log, srv, srv, srv, srv, srv, srv)
	api := &API{
		log:     log,
		cfg:     cfg,
		router:  mux.NewRouter(),
		db:      db,
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
		IdleTimeout:  api.cfg.IdleTimeout,
		WriteTimeout: api.cfg.IdleTimeout,
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
	public.HandleFunc("/getfavourites", api.handler.GetFavourites).Methods(http.MethodGet)
	public.HandleFunc("/addfavourite", api.handler.AddFavourite).Methods(http.MethodPost)
	public.HandleFunc("/removefavourite", api.handler.RemoveFavourite).Methods(http.MethodPost)
	public.HandleFunc("/getmale", api.handler.GetMale).Methods(http.MethodGet)
	public.HandleFunc("/getfemale", api.handler.GetFemale).Methods(http.MethodGet)
}
