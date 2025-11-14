package main

import (
	"paymentservice/internal/api"
	"paymentservice/internal/config"
	"paymentservice/internal/logger"
)

func main() {
	log := logger.InitLogger()
	cfg := config.InitConfig(log)
	api := api.InitAPI(log, cfg)
	api.StartServer()
}
