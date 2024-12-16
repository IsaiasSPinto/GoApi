package main

import (
	"github.com/IsaiasSPinto/GoApi/config"
	"github.com/IsaiasSPinto/GoApi/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("Error starting the application: %v", err)
		return
	}

	router.Init()

}
