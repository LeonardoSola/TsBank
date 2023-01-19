package routes

import (
	"net/http"

	controller "tsbank/controllers/auth"
)

// Rotas de autenticação
var AuthRoutes = []Route{
	// POST
	{
		URI:          "/auth/login",
		Method:       http.MethodPost,
		Handler:      controller.Login,
		AuthRequired: false,
	},
}
