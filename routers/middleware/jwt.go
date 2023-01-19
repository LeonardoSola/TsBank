package middleware

import (
	"net/http"
	"tsbank/auth"
	"tsbank/models"
	"tsbank/responses"

	"github.com/gin-gonic/gin"
)

// Authenticate é responsável por autenticar o usuário
func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if erro := auth.ValidateToken(ctx); erro != nil {
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
