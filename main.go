package main

import (
	"goodvandro/go-opportunities.git/config"
	"goodvandro/go-opportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errof("config init error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
