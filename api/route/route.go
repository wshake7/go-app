package route

import (
	"github.com/gin-gonic/gin"
	"go-app/bootstrap"
	"xorm.io/xorm"
)

func Setup(env *bootstrap.Env, engine *xorm.Engine, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewLoginRoute(env, engine, publicRouter)
}
