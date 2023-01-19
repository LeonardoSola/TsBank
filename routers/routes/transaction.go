package routes

import (
	"net/http"

	controller "tsbank/controllers/transaction"
)

var TransactionRoutes = []Route{
	// GET
	{
		URI:          "/transaction",
		Method:       http.MethodGet,
		Handler:      controller.Get,
		AuthRequired: true,
	},
	{
		URI:          "/transaction/{id}",
		Method:       http.MethodGet,
		Handler:      controller.GetByID,
		AuthRequired: true,
	},
	// POST
	{
		URI:          "/transaction",
		Method:       http.MethodPost,
		Handler:      controller.Create,
		AuthRequired: true,
	},
}
