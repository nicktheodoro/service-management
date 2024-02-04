package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var budgetRoutes = []Route{
	{
		URI:         "/budgets",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetBudgets,
		RequireAuth: true,
	},
	{
		URI:         "/budgets/:id",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetBudget,
		RequireAuth: true,
	},
	{
		URI:         "/budgets",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateBudget,
		RequireAuth: true,
	},
	{
		URI:         "/budgets/:id",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateBudget,
		RequireAuth: true,
	},
	{
		URI:         "/budgets/:id",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteBudget,
		RequireAuth: true,
	},
	{
		URI:         "/budgets/:id/update-status",
		Method:      http.MethodPatch,
		HandlerFunc: controllers.UpdateBudgetStatus,
		RequireAuth: true,
	},
}
