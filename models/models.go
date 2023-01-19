package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

func (p *Pagination) GetFromContext(ctx *gin.Context) {
	offsetQuery := ctx.Query("offset")
	limitQuery := ctx.Query("limit")

	limitInt, err := strconv.Atoi(limitQuery)
	if err != nil || limitInt == 0 {
		limitInt = 10
	}

	offsetInt, err := strconv.Atoi(offsetQuery)
	if err != nil || offsetInt == 0 {
		offsetInt = 0
	}

	p.Limit = limitInt
	p.Offset = offsetInt
}
