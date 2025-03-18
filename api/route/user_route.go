package route

import (
	"github.com/gin-gonic/gin"
	"go-app/api/controller"
	"go-app/bootstrap"
	"go-app/usecase"
)

func NewUserRoute(app *bootstrap.Application, group *gin.RouterGroup) {
	group = group.Group("user")
	con := controller.UserController{
		UserUseCase: usecase.UserUseCase{Engine: app.DBEngine},
		AclUseCase:  usecase.AclUseCase{Engine: app.DBEngine},
	}
	group.POST("/create", con.Create)
	group.GET("/http/acl/test", con.HttpAclTest)
}
