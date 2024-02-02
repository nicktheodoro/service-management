package test

import (
	"service-management/internal/pkg/config"
	"service-management/internal/pkg/database"
)

func SetupTests() {
	config.Setup("./config.yml")
	database.SetupDB()
	database.GetDB().Exec("DELETE FROM users")
	database.GetDB().Exec("DELETE FROM user_roles")
	database.GetDB().Exec("DELETE FROM customers")
	database.GetDB().Exec("DELETE FROM customers_address")
	database.GetDB().Exec("DELETE FROM services")
}
