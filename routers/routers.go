package routers

import (
	"tsbank/routers/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// SetupRoute é responsável por configurar as rotas da aplicação
func SetupRoute() *gin.Engine {
	enviroment := viper.GetBool("DEBUG")
	if enviroment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedIPs := viper.GetString("ALLOWED_IPS")

	router := gin.New()
	router.SetTrustedProxies([]string{allowedIPs})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	RegisterRoutes(router)

	return router
}
