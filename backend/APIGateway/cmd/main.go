package main

import (
	_ "apigateway/cmd/docs"
	"apigateway/internal/api"
	"apigateway/internal/config"
	"apigateway/internal/logger"
)

// @title MarketPlays API
// @version 1.0
// @description API Server for InternerShop App
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter: "Bearer {token}"
func main() {
	logger := logger.InitLogger()
	config := config.InitConfig(logger)
	api := api.InitAPI(logger, config)
	api.StartServer()
}
