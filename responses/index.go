package responses

import "github.com/gin-gonic/gin"

type response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func Data(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data: data,
	})
}
