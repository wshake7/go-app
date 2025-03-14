package route

import (
	"github.com/gin-gonic/gin"
	"go-app/bootstrap"
)

func Setup(app *bootstrap.Application, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewLoginRoute(app, publicRouter)
}
