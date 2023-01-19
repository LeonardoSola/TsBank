package responses

import (
	"tsbank/models"

	"github.com/gin-gonic/gin"
)

type response struct {
	Data       any                `json:"data,omitempty"`
	Error      string             `json:"error,omitempty"`
	Pagination *models.Pagination `json:"pagination,omitempty"`
}

func Data(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data: data,
	})
}

func List(ctx *gin.Context, status int, data interface{}, pagination models.Pagination) {
	ctx.JSON(status, response{
		Data:       data,
		Pagination: &pagination,
	})
}
