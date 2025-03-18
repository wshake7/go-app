package controller

import (
	"github.com/gin-gonic/gin"
	"go-app/domain"
	"go-app/usecase"
	"net/http"
)

type UserController struct {
	usecase.UserUseCase
	usecase.AclUseCase
}

func (receiver *UserController) Create(ctx *gin.Context) {
	receiver.AclUseCase.AclPermissionVerify(ctx, "user:create")
	ctx.JSON(http.StatusOK, domain.SuccessResponse())
}

func (receiver UserController) HttpAclTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, domain.SuccessResponse())
}
