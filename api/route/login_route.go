package route

import (
	"github.com/gin-gonic/gin"
	"go-app/api/controller"
	"go-app/bootstrap"
)

func NewLoginRoute(app *bootstrap.Application, group *gin.RouterGroup) {
	lc := controller.LoginController{
		Env:         app.Env,
		RedisClient: app.RedisClient,
	}
	group.GET("/login", lc.Login)
}
