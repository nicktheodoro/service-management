package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var customerAddressRoutes = []Route{
	{
		URI:         "/customers/:id/addresses",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetCustomerAddresses,
		RequireAuth: true,
	},
	{
		URI:         "/customers/:id/addresses/:addressId",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetCustomerAddress,
		RequireAuth: true,
	},
	{
		URI:         "/customers/:id/addresses",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateCustomerAddress,
		RequireAuth: true,
	},
	{
		URI:         "/customers/:id/addresses/:addressId",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateCustomerAddress,
		RequireAuth: true,
	},
	{
		URI:         "/customers/:id/addresses/:addressId",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteCustomerAddress,
		RequireAuth: true,
	},
}
