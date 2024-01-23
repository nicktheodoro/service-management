package router

import (
	"service-management/internal/api/middlewares"

	"github.com/gin-gonic/gin"
)

type Route struct {
	URI         string
	Method      string
	HandlerFunc gin.HandlerFunc
	RequireAuth bool
}

func Setup(mode string) *gin.Engine {
	app := gin.New()
	gin.SetMode(mode)
	return ConfigureRoutes(app)
}

// Configure adds all routes to the router.
func ConfigureRoutes(app *gin.Engine) *gin.Engine {
	routes := userRoutes
	routes = append(routes, authRoute)

	for _, route := range routes {
		if route.RequireAuth {
			app.Handle(route.Method, route.URI, middlewares.AuthRequired(), route.HandlerFunc)
		} else {
			app.Handle(route.Method, route.URI, route.HandlerFunc)
		}
	}

	return app
}
