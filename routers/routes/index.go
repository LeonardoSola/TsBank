package routes

import "github.com/gin-gonic/gin"

type Route struct {
	URI          string
	Method       string
	Handler      func(ctx *gin.Context)
	AuthRequired bool
}

// AllRoutes retorna a lista de todas as rotas da aplicação
func AllRoutes() (routes []Route) {
	routes = append(routes, AuthRoutes...)
	routes = append(routes, UserRoutes...)
	routes = append(routes, TransactionRoutes...)
	return
}
