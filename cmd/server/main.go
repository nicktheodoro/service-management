package main

import (
	_ "service-management/docs"
	"service-management/internal/api"
)

// @Golang API REST
// @version 1.0
// @description This API facilitates service management within a Golang application, leveraging the Gin Framework and Gorm ORM.

// @contact.name Nicolas Theodoro Lopes
// @contact.email nicolastheodoro@gmail.com

// @license.name MIT
// @license.url https://github.com/nicktheodoro/service-management/blob/main/LICENSE

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	api.Run("")
}
