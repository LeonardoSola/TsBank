package routers

import (
	"net/http"
	"tsbank/routers/middleware"
	"tsbank/routers/routes"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registra todas as rotas da aplicação
func RegisterRoutes(route *gin.Engine) {
	AllRoutes := routes.AllRoutes()

	api := route.Group("/api")
	//Add All route
	for _, r := range AllRoutes {
		if r.AuthRequired {
			api.Handle(r.Method, r.URI, middleware.Authenticate(), r.Handler)
		} else {
			api.Handle(r.Method, r.URI, r.Handler)
		}
	}

	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
}
