package test

import (
	"service-management/internal/pkg/config"
	"service-management/internal/pkg/database"
)

func SetupTests() {
	config.Setup("./config.yml")
	database.SetupDB()
	database.GetDB().Exec("DELETE FROM budgets")
	database.GetDB().Exec("DELETE FROM budget_items")
	database.GetDB().Exec("DELETE FROM customers")
	database.GetDB().Exec("DELETE FROM customers_address")
	database.GetDB().Exec("DELETE FROM services")
	database.GetDB().Exec("DELETE FROM providers")
	database.GetDB().Exec("DELETE FROM users")
	database.GetDB().Exec("DELETE FROM user_roles")
}
