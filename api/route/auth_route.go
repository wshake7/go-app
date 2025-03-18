package route

import (
	"github.com/gin-gonic/gin"
	"go-app/api/controller"
	"go-app/bootstrap"
	"go-app/usecase"
)

func NewAuthRoute(app *bootstrap.Application, group *gin.RouterGroup) {
	con := controller.AuthController{
		Env: app.Env,
		AuthUseCase: usecase.AuthUseCase{
			Engine: app.DBEngine,
		},
	}
	group.POST("/login", con.Login)
	group.POST("/signup", con.Signup)
	group.POST("/refresh", con.RefreshToken)
}
