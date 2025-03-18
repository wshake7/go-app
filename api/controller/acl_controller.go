package controller

import (
	"github.com/gin-gonic/gin"
	"go-app/bootstrap"
	"go-app/domain"
	"go-app/internal/utils"
	"go-app/usecase"
	"net/http"
)

type AclController struct {
	*bootstrap.Env
	*usecase.AclUseCase
}

func (receiver *AclController) PermissionCreate(ctx *gin.Context) {
	var registerPermissionRequest domain.PermissionCreateRequest
	err := ctx.ShouldBind(&registerPermissionRequest)
	utils.Panic(err)
	err = receiver.AclUseCase.CreatePermission(registerPermissionRequest.ToModel())
	utils.Panic(err)
	ctx.JSON(http.StatusOK, domain.SuccessResponse())
}

func (receiver *AclController) AclCreate(ctx *gin.Context) {
	var aclCreateRequest domain.AclCreateRequest
	err := ctx.ShouldBind(&aclCreateRequest)
	utils.Panic(err)
	acl := aclCreateRequest.ToModel()
	acl.UserId = ctx.GetInt64("userId")
	err = receiver.AclUseCase.CreateAcl(acl)
	utils.Panic(err)
	ctx.JSON(http.StatusOK, domain.SuccessResponse())
}
