package main

import (
	"CartService/internal/api"
	"CartService/internal/config"
	"CartService/internal/logger"
	"CartService/pkg"
	"os"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.InitConfig(logger)
	clientRedis, err := pkg.NewClient(logger, cfg)
	if err != nil {
		logger.Info("failed to create client redis.", "err", err)
		os.Exit(1)
	}
	api := api.InitApi(logger, cfg, clientRedis)
	api.StartServer()
}
