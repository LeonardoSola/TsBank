package middleware

import (
	"fmt"
	"net/http"
	"tsbank/auth"
	"tsbank/models"
	"tsbank/responses"

	"github.com/gin-gonic/gin"
)

// Authenticate é responsável por autenticar o usuário
func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Authenticating user", ctx.Request.Header.Get("Authorization"))
		if erro := auth.ValidateToken(ctx); erro != nil {
			fmt.Println("Erro ao validar token", erro)
			responses.Error(ctx, http.StatusUnauthorized, "Token não é valido")
			ctx.Abort()
			return
		}

		// get user and set to context
		userId, err := auth.ExtractUserID(ctx)
		if err != nil {
			responses.Error(ctx, http.StatusUnauthorized, "Token não é valido")
			ctx.Abort()
			return
		}

		var user models.User
		user.ID = userId
		user.FindById()

		ctx.Set("user", user)
		ctx.Next()
	}
}
