package router

import (
	"net/http"
	"service-management/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	var routes []Route
	routes = append(routes, authRoute)
	routes = append(routes, customerRoutes...)
	routes = append(routes, customerAddressRoutes...)
	routes = append(routes, userRoutes...)
	routes = append(routes, userRoleRoutes...)

	for _, route := range routes {
		if route.RequireAuth {
			app.Handle(route.Method, route.URI, middlewares.AuthRequired(), route.HandlerFunc)
		} else {
			app.Handle(route.Method, route.URI, route.HandlerFunc)
		}
	}

	app.Handle(http.MethodGet, "/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.NoRoute(middlewares.NoRouteHandler())
	return app
}
