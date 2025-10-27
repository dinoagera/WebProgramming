package api

import (
	"log/slog"
	"paymentservice/internal/config"

	"github.com/gorilla/mux"
)

type API struct {
	log    *slog.Logger
	cfg    *config.Config
	router *mux.Router
}

// func InitAPI(log *slog.Logger, cfg *config.Config) *API {

// }
