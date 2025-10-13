package api

import (
	"apigateway/internal/config"
	"log/slog"

	"github.com/gorilla/mux"
)

type API struct {
	log    *slog.Logger
	cfg    *config.Config
	router *mux.Router
	//Дописать слои для структуры API
}

// func InitAPI() *API {
// 	//
// }
