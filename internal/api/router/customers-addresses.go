package router

import (
	"net/http"
	"service-management/internal/api/controllers"
)

var customerAddressRoutes = []Route{
	{
		URI:         "/customers/addresses",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetCustomersAddresses,
		RequireAuth: true,
	},
	{
		URI:         "/customers/addresses/:id",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetCustomerAddress,
		RequireAuth: true,
	},
	// {
	// 	URI:         "/customers/addresses",
	// 	Method:      http.MethodPost,
	// 	HandlerFunc: controllers.CreateCustomerAddress,
	// 	RequireAuth: false,
	// },
	// {
	// 	URI:         "/customers/addresses/:id",
	// 	Method:      http.MethodPut,
	// 	HandlerFunc: controllers.UpdateCustomerAddress,
	// 	RequireAuth: true,
	// },
	// {
	// 	URI:         "/customers/addresses/:id",
	// 	Method:      http.MethodDelete,
	// 	HandlerFunc: controllers.DeleteCustomerAddress,
	// 	RequireAuth: true,
	// },
}
