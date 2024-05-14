package main

import (
	"github.com/Brendon-Marquezi/gopportunities/config"
	"github.com/Brendon-Marquezi/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}
