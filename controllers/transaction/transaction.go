package transaction

import (
	"net/http"
	"strconv"
	"tsbank/models"
	"tsbank/responses"

	"github.com/gin-gonic/gin"
)

// Get é responsável por retornar todas as transações de um usuário
func Get(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	var pagination models.Pagination
	pagination.GetFromContext(ctx)

	transaction := models.Transaction{}
	transactions, err := transaction.FindAll(user.ID, &pagination)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Erro ao listar transações")
		return
	}

	responses.List(ctx, http.StatusOK, transactions, pagination)
}

// GetByID é responsável por retornar uma transação
func GetByID(ctx *gin.Context) {

	transactionId := ctx.Param("id")

	transactionIdInt, err := strconv.Atoi(transactionId)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "ID inválido")
		return
	}

	transaction := models.Transaction{}
	transaction.ID = uint64(transactionIdInt)

	if err := transaction.FindById(); err != nil {
		responses.Error(ctx, http.StatusNotFound, "Transação não encontrada")
		return
	}

	userID := ctx.MustGet("user").(models.User).ID
	if transaction.OriginID != userID && transaction.DestinationID != userID {
		responses.Error(ctx, http.StatusForbidden, "Você não tem permissão para acessar essa transação")
		return
	}

	responses.Data(ctx, http.StatusOK, transaction)
}

// Create é responsável por criar uma transação
func Create(ctx *gin.Context) {
}
