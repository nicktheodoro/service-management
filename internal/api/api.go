package api

import (
	"fmt"
	"service-management/internal/api/router"
	"service-management/internal/pkg/config"
	"service-management/internal/pkg/database"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	database.SetupDB()
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data//config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup(conf.Server.Mode)
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
