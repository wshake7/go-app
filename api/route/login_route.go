package route

import (
	"github.com/gin-gonic/gin"
	"go-app/api/controller"
	"go-app/bootstrap"
	"go-app/usecase"
)

func NewLoginRoute(app *bootstrap.Application, group *gin.RouterGroup) {
	lc := controller.LoginController{
		Env:          app.Env,
		RedisClient:  app.RedisClient,
		LoginUseCase: usecase.NewLoginUseCase(app.DBEngine),
	}
	group.POST("/login", lc.Login)
	group.POST("/signup", lc.Signup)
	group.POST("/refresh", lc.RefreshToken)
	group.POST("/test", lc.Test)
}
