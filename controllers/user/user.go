package user

import (
	"net/http"
	"tsbank/models"
	"tsbank/responses"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	responses.Data(ctx, http.StatusOK, user)
}

// Create é responsável pelo cadastro de um novo usuário
func Create(ctx *gin.Context) {
	var user models.User

	// Lê os dados do usuário
	if err := ctx.ShouldBindJSON(&user); err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Cria o usuário
	if err := user.Create(); err != nil {
		responses.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Retorna o usuário criado
	responses.Data(ctx, http.StatusCreated, user)
}

func Deposit(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	var input struct {
		Value int64 `json:"value"`
	}

	if err := ctx.BindJSON(&input); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Dados inválidos")
		return
	}

	var valor int64 = int64(input.Value * 100)

	if err := user.Deposit(valor); err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	responses.Data(ctx, http.StatusOK, "Depósito realizado com sucesso")
}
