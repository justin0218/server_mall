package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"server_mall/configs"
	"server_mall/internal/controllers"
	//"server_mall/internal/middleware"
	//
	//"server_mall/internal/routers/v1/ws"
)

func Init() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.Dft.Get().Runmode)
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "*",
		ExposedHeaders:  "*",
		Credentials:     true,
		ValidateHeaders: false,
	}))
	global := r.Group("mall")
	global.GET("health", func(context *gin.Context) {
		context.JSON(200, map[string]string{"msg": "ok"})
		return
	})
	auth := new(controllers.AuthController)
	global.GET("/v1/client/login", auth.Login)
	global.GET("/v1/server/access_token", auth.AccessToken)
	return r
}
