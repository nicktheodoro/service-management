package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var customerRoutes = []Route{
	{
		URI:         "/customers",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetCustomers,
		RequireAuth: true,
	},
	{
		URI:         "/customers/:id",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetCustomer,
		RequireAuth: true,
	},
	{
		URI:         "/customers",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateCustomer,
		RequireAuth: true,
	},
	{
		URI:         "/customers/:id",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateCustomer,
		RequireAuth: true,
	},
	{
		URI:         "/customers/:id",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteCustomer,
		RequireAuth: true,
	},
}
