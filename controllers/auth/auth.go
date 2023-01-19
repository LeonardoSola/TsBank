package auth

import (
	"net/http"
	"strings"
	"tsbank/auth"
	"tsbank/models"
	"tsbank/responses"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.User

	// Lê os dados do usuário
	if err := ctx.BindJSON(&user); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Erro ao ler os dados do usuário")
		return
	}

	// Verifica dados obrigatórios
	if user.Email == "" || user.Pass == "" {
		responses.Error(ctx, http.StatusBadRequest, "Usuario ou senha inválidos")
		return
	}

	// Verifica se o usuário existe
	var userDB models.User
	userDB.Email = strings.ToLower(user.Email)

	err := userDB.FindByEmail()
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Usuario ou senha inválidos")
		return
	}

	// Verifica se a senha está correta
	if !userDB.CheckPassword(user.Pass) {
		responses.Error(ctx, http.StatusUnauthorized, "Usuario ou senha inválidos")
		return
	}

	// Gera o token
	token, err := auth.GenToken(userDB.ID)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Erro ao gerar o token")
		return
	}

	// Retorna o token
	responses.Data(ctx, http.StatusOK, token)
}
