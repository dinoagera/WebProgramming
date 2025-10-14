package main

import (
	"apigateway/internal/api"
	"apigateway/internal/config"
	"apigateway/internal/logger"
)

func main() {
	logger := logger.InitLogger()
	config := config.InitConfig(logger)
	api := api.InitAPI(logger, config)
	api.StartServer()
}
