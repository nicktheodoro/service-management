package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var userRoleRoutes = []Route{
	{
		URI:         "/users/roles",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUserRoles,
		RequireAuth: true,
	},
	{
		URI:         "/users/roles/:id",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUserRole,
		RequireAuth: true,
	},
	{
		URI:         "/users/roles",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateUserRole,
		RequireAuth: true,
	},
	{
		URI:         "/users/roles/:id",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateUserRole,
		RequireAuth: true,
	},
	{
		URI:         "/users/roles/:id",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteUserRole,
		RequireAuth: true,
	},
}
