package routes

import (
	"net/http"

	controller "tsbank/controllers/user"
)

var UserRoutes = []Route{
	// GET
	{
		URI:          "/user",
		Method:       http.MethodGet,
		Handler:      controller.Get,
		AuthRequired: true,
	},
	// POST
	{
		URI:          "/user",
		Method:       http.MethodPost,
		Handler:      controller.Create,
		AuthRequired: false,
	},
	// PUT
	{
		URI:          "/user/deposit",
		Method:       http.MethodPut,
		Handler:      controller.Deposit,
		AuthRequired: true,
	},
}
