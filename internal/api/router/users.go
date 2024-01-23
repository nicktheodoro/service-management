package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUsers,
		RequireAuth: true,
	},
	{
		URI:         "/users/:id",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUser,
		RequireAuth: true,
	},
	{
		URI:         "/users",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/:id",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/:id",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteUser,
		RequireAuth: true,
	},
}
