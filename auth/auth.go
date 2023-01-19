package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var ExpireTime = func() int64 {
	return time.Now().Add(time.Hour * 24).Unix()
}

func GenToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{
		"exp":    ExpireTime(),
		"userId": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(viper.GetString("SECRET_KEY")))
}

func ValidateToken(ctx *gin.Context) error {
	tokenString := extractToken(ctx)

	token, err := jwt.Parse(tokenString, verificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalido")
}

func extractToken(ctx *gin.Context) string {
	token := ctx.Request.Header["Authorization"]

	if len(token) == 0 {
		return ""
	}

	if len(strings.Split(token[0], " ")) == 2 {
		return strings.Split(token[0], " ")[1]
	}

	return ""
}

func verificationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperado! %v", token.Header["alg"])
	}

	return []byte(viper.GetString("SECRET_KEY")), nil
}

func ExtractUserID(ctx *gin.Context) (userID uint64, err error) {
	tokenString := extractToken(ctx)

	token, err := jwt.Parse(tokenString, verificationKey)
	if err != nil {
		return
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID64, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID64, nil
	}

	return 0, fmt.Errorf("token invalido")
}
