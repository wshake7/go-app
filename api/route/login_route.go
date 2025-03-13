package route

import (
	"github.com/gin-gonic/gin"
	"go-app/api/controller"
	"go-app/bootstrap"
	"xorm.io/xorm"
)

func NewLoginRoute(env *bootstrap.Env, engine *xorm.Engine, group *gin.RouterGroup) {
	lc := controller.LoginController{
		Env: env,
	}
	group.POST("/login", lc.Login)
}
