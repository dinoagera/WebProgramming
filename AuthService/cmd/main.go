package main

import (
	"authservice/internal/api"
	"authservice/internal/config"
	"authservice/internal/logger"

	_ "authservice/cmd/docs"
)

// @title Auth API
// @version 1.0
// @description API Server for InternerShop App

// @host localhost:8081
// @BasePath /api
func main() {
	log := logger.InitLogger()
	cfg := config.InitConfig(log)
	api := api.InitApi(log, cfg)
	api.StartServer()
}
