package main

import (
	"catalogservice/internal/api"
	"catalogservice/internal/config"
	"catalogservice/internal/logger"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.InitConfig(logger)
	api := api.InitAPI(logger, cfg)
	api.StartServer()
}
