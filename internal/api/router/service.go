package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var serviceRoutes = []Route{
	{
		URI:         "/services",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetServices,
		RequireAuth: true,
	},
	{
		URI:         "/services/:id",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetService,
		RequireAuth: true,
	},
	{
		URI:         "/services",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateService,
		RequireAuth: true,
	},
	{
		URI:         "/services/:id",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateService,
		RequireAuth: true,
	},
	{
		URI:         "/services/:id",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteService,
		RequireAuth: true,
	},
}
