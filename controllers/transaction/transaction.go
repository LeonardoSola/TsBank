package transaction

import (
	"net/http"
	"strconv"
	"tsbank/models"
	"tsbank/responses"
	"tsbank/security"

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
	user := ctx.MustGet("user").(models.User)

	if user.CanTransfer() {
		responses.Error(ctx, http.StatusForbidden, "Você não pode realizar transferências")
		return
	}

	var input models.Transaction

	if err := ctx.BindJSON(&input); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Dados inválidos")
		return
	}

	transaction := models.Transaction{
		OriginID:      user.ID,
		DestinationID: input.DestinationID,
		Value:         input.Value,
	}

	destiny := models.User{}
	destiny.ID = transaction.DestinationID
	if err := destiny.FindById(); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Conta de destino não encontrada")
		return
	}

	if err := user.Withdraw(input.Value); err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transaction.Authorized = security.AuthorizeTransaction()

	if err := transaction.Create(); err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !transaction.Authorized {
		responses.Error(ctx, http.StatusServiceUnavailable, "Transação não autorizada")
		user.Deposit(input.Value)
		return
	} else {
		destiny.Deposit(input.Value)
	}

	responses.Data(ctx, http.StatusCreated, transaction)
}
