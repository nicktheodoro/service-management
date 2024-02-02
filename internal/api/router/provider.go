package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var providerRoutes = []Route{
	{
		URI:         "/providers",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetProviders,
		RequireAuth: true,
	},
	{
		URI:         "/providers/:id",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetProvider,
		RequireAuth: true,
	},
	{
		URI:         "/providers",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateProvider,
		RequireAuth: true,
	},
	{
		URI:         "/providers/:id",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateProvider,
		RequireAuth: true,
	},
	{
		URI:         "/providers/:id",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteProvider,
		RequireAuth: true,
	},
}
