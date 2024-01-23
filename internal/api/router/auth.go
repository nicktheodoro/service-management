package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var authRoute = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	HandlerFunc: controllers.Login,
	RequireAuth: false,
}
