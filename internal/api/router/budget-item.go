package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var budgetItemRoutes = []Route{
	{
		URI:         "/budgets/:id/items",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetBudgetItems,
		RequireAuth: true,
	},
	{
		URI:         "/budgets/:id/items/:itemId",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetBudgetItem,
		RequireAuth: true,
	},
	{
		URI:         "/budgets/:id/items",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateBudgetItem,
		RequireAuth: true,
	},
	{
		URI:         "/budgets/:id/items/:itemId",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateBudgetItem,
		RequireAuth: true,
	},
	{
		URI:         "/budgets/:id/items/:itemId",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteBudgetItem,
		RequireAuth: true,
	},
}
