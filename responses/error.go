package responses

import "github.com/gin-gonic/gin"

// Error é responsável por retornar os erros da aplicação
func Error(ctx *gin.Context, code int, err string) {
	ctx.JSON(code, response{
		Error: err,
	})
}
