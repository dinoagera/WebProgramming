package main

import (
	"apigateway/internal/config"
	"apigateway/internal/logger"
)

func main() {
	logger := logger.InitLogger()
	config := config.InitConfig(logger)
}
