package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var userRoleRoutes = []Route{
	{
		URI:         "/user-roles",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUserRoles,
		RequireAuth: true,
	},
	{
		URI:         "/user-roles/:id",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUserRole,
		RequireAuth: true,
	},
	{
		URI:         "/user-roles",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateUserRole,
		RequireAuth: true,
	},
	{
		URI:         "/user-roles/:id",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateUserRole,
		RequireAuth: true,
	},
	{
		URI:         "/user-roles/:id",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteUserRole,
		RequireAuth: true,
	},
}
