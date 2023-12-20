package main

import (
	"goodvandro/go-opportunities/config"
	"goodvandro/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
