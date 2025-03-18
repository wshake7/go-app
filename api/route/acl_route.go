package route

import (
	"github.com/gin-gonic/gin"
	"go-app/api/controller"
	"go-app/bootstrap"
	"go-app/usecase"
)

func NewAclRoute(app *bootstrap.Application, aclUseCase *usecase.AclUseCase, group *gin.RouterGroup) {
	con := controller.AclController{
		Env:        app.Env,
		AclUseCase: aclUseCase,
	}
	group.POST("/permission/create", con.PermissionCreate)
	group.POST("/acl/create", con.AclCreate)
}
